package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
)

//可以收集多个日志
//exitChan:当删除一个任务的时候,需要通知退出
type TailObj struct {
	tail     *tail.Tail
	secLimit *SecondLimit
	//每个tailf记录偏移量
	offset int64
	//对应的文件
	logConf  LogConfig
	exitChan chan bool
}

//使用map存储,可以去重
type TailMgr struct {
	tailObjMap map[string]*TailObj
	lock       sync.Mutex
}

//多少文件就需要多少tailf
var tailMgr *TailMgr

func RunServer() {
	//1.创建tailf实例
	tailMgr = NewTailMgr()
	//2.拿到处理后的消息
	tailMgr.process()
	waitGroup.Wait()
}

func NewTailMgr() (tailMgr *TailMgr) {
	tailMgr = &TailMgr{
		tailObjMap: make(map[string]*TailObj, 8),
	}
	return
}

func (t *TailMgr) process() {
	//从管道拿到消息,反序列化成logconfig
	var logConfArr []LogConfig
	for v := range GetLogConfChan() {
		err := json.Unmarshal([]byte(v), logConfArr)
		if err != nil {
			continue
		}
		//管道拿到消息,一个json字符串,去解析
		err = t.reloadConfig(logConfArr)
		if err != nil {
			continue
		}
	}
}

//拿到etcd最新的配置,去热加载
//tailMgr管理的是正在运行的所有tail实例
//如果存在就更新,不存在需要添加一个新的tail去收集
func (t *TailMgr) reloadConfig(logConfArr []LogConfig) (err error) {
	//处理新增 修改任务
	for _, conf := range logConfArr {
		tailObj, ok := t.tailObjMap[conf.LogPath]
		if !ok {
			err := t.AddLogFile(conf)
			if err != nil {
				continue
			}
			continue
		}
		tailObj.logConf = conf
		tailObj.secLimit.limit = int32(conf.SendRate)
		t.tailObjMap[conf.LogPath] = tailObj
	}

	//处理 删除任务
	//当前正在跑的任务,如果不在最新配置中,就删除
	for key, tailObj := range t.tailObjMap {
		var found = false
		for _, newVale := range logConfArr {
			if key == newVale.LogPath {
				found = true
				break
			}
		}
		if found == false {
			//通知readlog读取日志线程退出
			tailObj.exitChan <- true
			delete(t.tailObjMap, key)
		}
	}
	return
}

//二次进行校验,防止有重复日志收集
func (t *TailMgr) AddLogFile(conf LogConfig) (err error) {
	t.lock.Lock()
	defer t.lock.Unlock()
	//如果存在,就不添加 避免重复收集
	_, ok := t.tailObjMap[conf.LogPath]
	if ok {
		err = fmt.Errorf("duplicate filename:%s", conf.LogPath)
		return
	}
	//如果不存在,就初始化一个tail实例
	tail, err := tail.TailFile(conf.LogPath, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从哪个位置读取
		MustExist: false,
		Poll:      true, //轮循
	})

	tailObj := &TailObj{
		secLimit: NewSecondLimit(int32(conf.SendRate)),
		logConf:  conf,
		offset:   0,
		tail:     tail,
		exitChan: make(chan bool, 1),
	}
	t.tailObjMap[conf.LogPath] = tailObj
	waitGroup.Add(1)
	//通过tail
	go tailObj.readLog()
	return
}

func (t *TailObj) readLog() {
	//获取 从tail读取的数据
	for line := range t.tail.Lines {
		if line.Err != nil {
			logs.Error("read line failed, err:%v", line.Err)
			continue
		}

		str := strings.TrimSpace(line.Text)
		if len(str) == 0 || str[0] == '\n' {
			continue
		}
		//读取到数据加到kafka中  kafka也是单独一组线程
		//放入kafkachan中,kafka有一个线程在监听,有数据就发送
		kafkaSender.addMessage(line.Text, t.logConf.Topic)

		//发送kafka同时,判断是不是达到上线
		t.secLimit.Add(1)
		t.secLimit.Wait()

		//判断退出 是不是删除了任务,如果是就退出当前任务,不再读取
		select {
		case <-t.exitChan:
			logs.Warn("tail obj is exited, config:%v", t.logConf)
			return
		default:
		}
	}
	waitGroup.Done()
}

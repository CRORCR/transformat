package main

import (
	"encoding/json"
	"strings"
	"sync"

	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
)

//可以收集多个日志
type TailObj struct {
	tail     *tail.Tail
	secLimit *SecondLimit
	offset   int64

	logConf  LogConfig
	exitChan chan bool
}

type TailMgr struct {
	tailObjMap map[string]*TailObj
	lock       sync.Mutex
}

var tailMgr *TailMgr

func RunServer() {
	tailMgr = NewTailMgr()
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

func (t *TailMgr) reloadConfig(logConfArr []LogConfig) (err error) {
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

	//处理热加载
	for key, tailObj := range t.tailObjMap {
		var found = false
		for _, newVale := range logConfArr {
			if key == newVale.LogPath {
				found = true
				break
			}
		}
		if found == false {
			tailObj.exitChan <- true
			delete(t.tailObjMap, key)
		}
	}
	return
}

func (t *TailMgr) AddLogFile(conf LogConfig) (err error) {
	t.lock.Lock()
	defer t.lock.Unlock()
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
		//数据加到kafka中
		kafkaSender.addMessage(line.Text, t.logConf.Topic)
		t.secLimit.Add(1)
		t.secLimit.Wait()

		//判断退出
		select {
		case <-t.exitChan:
			logs.Warn("tail obj is exited, config:%v", t.logConf)
			return
		default:
		}
	}
	waitGroup.Done()
}

package model

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
)

var (
	pool *redis.Pool
)


//init方法就是初始化redis etcd等等配置的方法
//可以传入对应的字段,但是后期如果需要加载其他的配置不好维护
//所以这里使用结构体,传入一个结构体,比较好维护
type ModelConf struct{
	RedisAddr string
	RedisPasswd string

	EtcdAddr       string
	EtcdProductKey string

	SendQueueName string
	RecvQueueName string

	SendQueueThreadNum int
	RecvQueueThreadNum int

	LogPath  string
	LogLevel string
}

func Init(conf *ModelConf)(err error){
	//1.初始化redis
	err = initRedis(conf)
	if err!=nil{
		return
	}
	//2.加载etcd(第一次)
	err = initEtcd(conf)
	if err != nil {
		return
	}
	//3.加载最新etcd配置(热加载)(从管道中获得商品数据,数据是从etcd获得的,然后存入的管道中)
	err = loadProductInfo()
	if err != nil {
		return
	}
	//4.起两组线程 接收和发送
	err=initRecvThread(conf)
	if err!=nil{
		return
	}
	err=initSendThread(conf)
	if err!=nil{
		return
	}
	return
}

func initRedis(conf *ModelConf)(err error){
	//1.创建线程池,第二个参数密码可选
	pool=newPool(conf.RedisAddr,conf.RedisPasswd)
	//2.获得连接
	conn:=pool.Get()
	//3.ping一下测试连接
	_,err=conn.Do("ping")
	if err!=nil{
		logs.Error("connect to redis failed, err:%v", err)
		return
	}
	logs.Debug("connect to redis succ")
	return
}

//拿到商品id,获取商品信息
func SecInfo(product_id int)(a *Activity,err error){
	a,err=secProxyData.GetActive(product_id)
	return
}

//从管道获取
func loadProductInfo()(err error){
	//1.初始化结构
	secProxyData=&SecProxyData{
		activityMap:make(map[int]*Activity,128),
		requestChan:make(chan *SecKillRequest,1000),
	}
	//2.从管道获取(etcd中是json字符串)
	productConf:=<-GetProductChan()
	//3.反序列化  活动最新配置就在activeArry中
	var activeArry []*Activity
	err=json.Unmarshal([]byte(productConf),&activeArry)
	if err!=nil{
		return
	}
	//3.活动加到map中,同时更新map(第一次)
	err=secProxyData.UpdateActive(activeArry)
	if err!=nil{
		return
	}
	//4.启动线程去热加载
	go secProxyData.Reload()
	return
}

func SecKill(productId,userId int,userIp string)(result *SecKillResult,err error){
	//1.生成秒杀请求 一个请求对应一个chan,长度就为1
	req := &SecKillRequest{
		ProductId:  productId,
		UserId:     userId,
		UserIp:     userIp,
		resultChan: make(chan *SecKillResult, 1),
	}
	//2.请求放入chan中,发送线程给redis发送
	err = secProxyData.AddRequest(req)
	if err != nil {
		logs.Error("add request failed, err:%v", err)
		return
	}
	//3.发送成功,等待结果,也需要设置超时
	timer := time.NewTicker(10 * time.Second)
	select {
	case <-timer.C:
		err = fmt.Errorf("timeout")
	case result = <-req.resultChan:
		return
	}
	return
}









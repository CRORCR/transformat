package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"transformat/seckill/proxy/model"
)

//定义一个结构,接收config配置文件数据
var proxyConf model.ModelConf

func initConf() (err error) {
	redisAddr := beego.AppConfig.String("redis::redis_addr")
	if len(redisAddr) == 0 {
		err = fmt.Errorf("invalid redis addr")
		return
	}
	sendQueue := beego.AppConfig.String("redis::send_request_queue_name")
	if len(sendQueue) == 0 {
		err = fmt.Errorf("invalid redis sendQueue")
		return
	}
	recvQueue := beego.AppConfig.String("redis::send_request_queue_name")
	if len(recvQueue) == 0 {
		err = fmt.Errorf("invalid redis recvQueue")
		return
	}
	sendNum,err := beego.AppConfig.Int("redis::send_queue_thread_num")
	if err!=nil {
		err = fmt.Errorf("invalid redis recvQueue")
		return
	}
	recNum,err := beego.AppConfig.Int("redis::recv_queue_thread_num")
	if err!=nil {
		err = fmt.Errorf("invalid redis recvQueue")
		return
	}
	//读取两个日志配置
	logPath := beego.AppConfig.String("logs::log_path")
	if len(logPath) == 0 {
		err = fmt.Errorf("invalid logs logPath")
		return
	}
	logLevel := beego.AppConfig.String("logs::log_level")
	if len(logLevel) == 0 {
		err = fmt.Errorf("invalid logs logLevel")
		return
	}
	etcdAddr := beego.AppConfig.String("etcd::addr")
	if (len(etcdAddr) == 0) {
		err = fmt.Errorf("invalid etcdAddr")
		return
	}

	etcdProductKey := beego.AppConfig.String("etcd::product_key")
	if (len(etcdProductKey) == 0) {
		err = fmt.Errorf("invalid etcdProductKey")
		return
	}
	proxyConf.RedisAddr = redisAddr
	proxyConf.SendQueueName = sendQueue
	proxyConf.RecvQueueName = recvQueue
	proxyConf.SendQueueThreadNum=sendNum
	proxyConf.RecvQueueThreadNum=recNum

	proxyConf.LogLevel = logLevel
	proxyConf.LogPath = logPath

	proxyConf.EtcdAddr=etcdAddr
	proxyConf.EtcdProductKey=etcdProductKey
	return
}

package main

import "github.com/astaxie/beego/logs"

func main() {
	err := initConfig("./conf/app.conf")
	if err != nil {
		logs.Error("inint config failed")
		panic(err)
	}
	initLog()
	initKafka()
	err = getLocalIP()
	if err != nil {
		logs.Error("get local ip failed, er:%v", err)
		return
	}
	err = initEtcd()
	if err != nil {
		logs.Error("init etcd failed")
		return
	}
	err = initES(appConfig.esAddr)
	if err != nil {
		logs.Error("init es failed, err:%v", err)
		return
	}
	err = Run(appConfig.esThreadNum)
	if err != nil {
		logs.Error("run es failed, err:%v", err)
		return
	}
	logs.Debug("run exited")
}

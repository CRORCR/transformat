package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/astaxie/beego/logs"
)

func main() {
	//1.初始化配置文件
	err := initConfig("./conf/app.conf")
	if err != nil {
		panic(fmt.Sprintf("init config failed,err:%d\n", err))
	}
	//2.初始化log
	initLog()

	//获得ip
	err = getIp()
	if err != nil {
		panic(err)
	}
	//3.初始化kafka
	initKafka()
	//4.初始化etcd
	err = initEtcd()
	if err != nil {
		logs.Error("init etcd failed, err:%v", err)
		return
	}
	RunServer()
}

package main

import (
	"fmt"
	"transformat/seckill/proxy/model"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "transformat/seckill/proxy/route"
)

func main() {
	//1.初始化redis配置
	err := initConf()
	if err != nil {
		//如果配置都出错,就直接panic
		panic(fmt.Sprintf("init config failed,err:%v", err))
	}

	//2.初始化logs
	initLogs(proxyConf.LogPath, proxyConf.LogLevel)
	//2.1 始化成功logs,就可以使用logs往文件输出日志了
	logs.Debug("init config and log success")
	//%#v 输出json格式
	logs.Debug("配置文件值为:%#v",proxyConf)

	//3.初始化所有服务
	err=initModel()
	if err!=nil{
		logs.Error("init model failed, err:%v", err)
		return
	}

	beego.Run()
}

//给所有服务初始化
//给指针,避免大数据的拷贝
func initModel()(err error){
	err= model.Init(&proxyConf)
	return
}


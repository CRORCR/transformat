package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

//读取配置库

func main() {
	conf,err:=config.NewConfig("ini","./logcollect.conf")
	if err!=nil{
		fmt.Printf("new config is failed ,err:%v \n",err)
		return
	}
	port,err:=conf.Int("server::port")
	if err!=nil{
		fmt.Printf("read port is failed,err:%v \n",err)
		return
	}
	fmt.Println("port:",port)

	level,err:=conf.Int("log::log_level")
	if err!=nil{
		fmt.Printf("read log_level is failed,err:%v \n",err)
		return
	}
	fmt.Println("level:",level)

	path:=conf.String("log::log_path")
	fmt.Println("path:",path)
}

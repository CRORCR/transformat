package main

import (
	"time"
	"transformat/config"

	"fmt"
	"sync/atomic"
)

type AppConfig struct {
	port      int
	kafkaAddr string
}

//定义一个atomic.value
type AppConfigMgr struct {
	config atomic.Value
}

var appConfigMgr = &AppConfigMgr{}

func main() {
	conf, err := config.NewConfig("./config.conf")
	if err != nil {
		fmt.Println("parse config failed")
		return
	}
	//传递的是原子value
	conf.AddNotifyer(appConfigMgr)
	//存储具体的配置结构
	var appConfig = &AppConfig{}
	appConfig.port, err = conf.GetInt("server_port")
	if err != nil {
		fmt.Println("get port failed, err:", err)
		return
	}

	fmt.Println("port:", appConfig.port)

	appConfig.kafkaAddr, err = conf.GetString("kafka_addr")
	if err != nil {
		fmt.Println("get kafkaAddr failed, err:", err)
		return
	}

	fmt.Println("kafkaAddr:", appConfig.kafkaAddr)
	//配置放到value中,后面使用就不用加锁了
	appConfigMgr.config.Store(appConfig)

	run()
}

//使用原子操作的value去实现接口
func (a *AppConfigMgr) Callback(conf *config.Config) {
	var appConfig = &AppConfig{}

	port, err := conf.GetInt("server_port")
	if err != nil {
		fmt.Println("get port failed, err:", err)
		return
	}

	appConfig.port = port
	fmt.Println("port:", appConfig.port)

	kafkaAddr, err := conf.GetString("kafka_addr")
	if err != nil {
		fmt.Println("get kafkaAddr failed, err:", err)
		return
	}

	appConfig.kafkaAddr = kafkaAddr
	fmt.Println("kafkaAddr:", appConfig.kafkaAddr)
	//把value中的旧配置替换掉,不用担心是否正在使用,atomic是原子的
	appConfigMgr.config.Store(appConfig)
}

func run() {
	for {
		//从value读取出来,强转成appconfig,然后就可以读操作
		appConfig := appConfigMgr.config.Load().(*AppConfig)
		fmt.Println("port:", appConfig.port)
		fmt.Println("kafkaAddr:", appConfig.kafkaAddr)
		fmt.Println()
		fmt.Println()
		time.Sleep(5 * time.Second)
	}
}

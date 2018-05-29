package main

import (
	"fmt"
	"strings"
	"transformat/config"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

//定义一个结构,接收配置文件的参数
type AppConfig struct {
	LogPath        string
	LogLevel       string
	kafkaAddr      string
	KafkaThreadNum int
	EtcdAddr       []string
	//单位为毫秒
	etcdTimeout     int
	etcdWatchKeyFmt string
}

//成员变量,赋值
var appConfig = &AppConfig{}

func initConfig(filename string) (err error) {
	conf, err := config.NewConfig(filename)
	if err != nil {
		return
	}

	logPath, err := conf.GetString("log_path")
	if err != nil || len(logPath) == 0 {
		return
	}

	logLevel, err := conf.GetString("log_level")
	if err != nil || len(logPath) == 0 {
		return
	}

	kafkaAddr, err := conf.GetString("kafka_addr")
	if err != nil || len(logPath) == 0 {
		return
	}
	etcdAddr, err := conf.GetString("etcd_addr")
	if err != nil || len(logPath) == 0 {
		return
	}

	arr := strings.Split(etcdAddr, ",")
	for _, v := range arr {
		str := strings.TrimSpace(v)
		if len(str) == 0 {
			continue
		}
		appConfig.EtcdAddr = append(appConfig.EtcdAddr, str)
	}
	appConfig.etcdTimeout = conf.GetIntDefault("etcd_timeout", 1500)
	appConfig.KafkaThreadNum = conf.GetIntDefault("kafka_thread_num", 8)
	etcdKey, err := conf.GetString("etcd_watch_key")
	if err != nil || len(etcdKey) == 0 {
		logs.Warn("get etcd watch key failed, err:%v", err)
		return
	}

	appConfig.etcdWatchKeyFmt = etcdKey
	appConfig.kafkaAddr = kafkaAddr
	appConfig.LogLevel = logLevel
	appConfig.LogPath = logPath

	fmt.Printf("load config succ, data:%v\n", appConfig)
	return
}

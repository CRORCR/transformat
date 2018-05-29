package main

import (
	"encoding/json"
	"fmt"
	"transformat/config"

	"github.com/astaxie/beego/logs"
)

type AppConfig struct {
	logPath       string
	logLevel      string
	kafkaAddr     string
	esAddr        string
	esThreadNum   int
	etcdAddr      string
	etcdKeyFormat string
}

var appConfig AppConfig

func initConfig(fileName string) (err error) {

	conf, err := config.NewConfig(fileName)
	if err != nil {
		return
	}

	logPath, err := conf.GetString("log_path")
	if len(logPath) == 0 || err != nil {
		return fmt.Errorf("get log_path failed,invalid logPath, err:%v", err)
	}

	appConfig.logPath = logPath
	logLevel, err := conf.GetString("log_level")
	if len(logLevel) == 0 || err != nil {
		return fmt.Errorf("get logLevel failed,invalid logLevel, err:%v", err)
	}

	appConfig.logLevel = logLevel

	kafkaAddr, err := conf.GetString("kafka_addr")
	if len(kafkaAddr) == 0 || err != nil {
		return fmt.Errorf("get kafkaAddr failed,invalid kafkaAddr, err:%v", err)
	}

	appConfig.kafkaAddr = kafkaAddr

	esAddr, err := conf.GetString("es_addr")
	if len(kafkaAddr) == 0 || err != nil {
		return fmt.Errorf("get es_addr failed,invalid es_addr, err:%v", err)
	}

	appConfig.esAddr = esAddr

	esThreadNum := conf.GetIntDefault("es_thread_num", 8)
	appConfig.esThreadNum = esThreadNum

	etcdAddr, err := conf.GetString("etcd_addr")
	if len(etcdAddr) == 0 || err != nil {
		return fmt.Errorf("get etcdAddr failed,invalid etcdAddr, err:%v", err)
	}

	appConfig.etcdAddr = etcdAddr

	etcdKey, err := conf.GetString("etcd_transfer_key")
	if len(etcdAddr) == 0 || err != nil {
		return fmt.Errorf("get etcd_transfer_key failed,invalid etcd_transfer_key, err:%v", err)
	}
	appConfig.etcdKeyFormat = etcdKey
	return
}

func initLog() {
	//从配置文件中读取到log地址和级别
	m := make(map[string]interface{})
	m["filename"] = appConfig.logPath
	m["level"] = getLevel(appConfig.logLevel)
	data, err := json.Marshal(m)
	logs.SetLogger(logs.AdapterFile, string(data))
	if err != nil {
		logs.Error("init log failed")
		panic("init log failed")
	}
}

func getLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "trace":
		return logs.LevelTrace
	case "warn":
		return logs.LevelWarning
	case "info":
		return logs.LevelInformational
	case "error":
		return logs.LevelError
	default:
		return logs.LevelDebug
	}
}

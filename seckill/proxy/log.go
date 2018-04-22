package main

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

func GetLevel(level string) int {
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

func initLogs(logPath, logLevel string) (err error) {
	config := make(map[string]interface{})
	config["fileName"] = logPath
	config["level"] = GetLevel(logLevel)
	//序列化
	data, err := json.Marshal(config)
	if err != nil {
		return
	}
	//初始化beego日志 第一个参数固定,输出到文件 第二个参数json
	logs.SetLogger(logs.AdapterFile, string(data))
	return
}

package main

import (
	"encoding/json"

	"github.com/astaxie/beego/logs"
)

func initLog() {
	m := make(map[string]interface{})
	m["filename"] = appConfig.LogPath
	m["level"] = getLevel(appConfig.LogLevel)
	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	logs.SetLogger(logs.AdapterFile, string(data))
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

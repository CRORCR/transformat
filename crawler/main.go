package main

import (
	"transformat/crawler/engine"
	"transformat/crawler/zhenai/parse"
)

func main() {
	//种子消息发送给engine 请求和解析器
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parse.ParseCityList,
	})
}

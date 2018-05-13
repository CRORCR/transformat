package main

import (
	"transformat/crawler/engine"
	"transformat/crawler/zhenai/parse"
)

func main() {
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parse.ParseCityList,
	})
}

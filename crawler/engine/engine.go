package engine

import (
	"log"
	"transformat/crawler/fetcher"
)

//可以接收很多种子地址
func Run(send ...Request) {
	var requests []Request
	requests = append(requests, send...)

	//循环读取request,爬取信息
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		//1.调用fetcher 返回所有utf8数据
		body, err := fetcher.Fetcher(r.Url)
		if err != nil {
			//有错,忽略
			log.Printf("fetcher:error url:%s err:%v", r.Url, err)
			continue
		}
		//2.返回文件给parse去解析,返回结果(请求和url)
		parseResult := r.ParseFunc(body)
		//3.请求放入消息列表
		requests = append(requests, parseResult.Request...)
		//item先输出 不管
		for _, v := range parseResult.Item {
			log.Printf("got item:%v \n", v)
		}
	}
}

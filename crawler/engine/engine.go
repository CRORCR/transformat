package engine

import (
	"log"
	"transformat/crawler/fetcher"
)

//可以接收很多种子地址
func Run(send ...Request) {
	var requests []Request
	requests = append(requests, send...)

	//只要有请求就爬取
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		//返回所有utf8数据
		body, err := fetcher.Fetcher(r.Url)
		if err != nil {
			//有错,忽略
			log.Printf("fetcher:error url:%s err:%v", r.Url, err)
			continue
		}
		parseResult := r.ParseFunc(body)
		requests = append(requests, parseResult.Request...)

		for _, v := range parseResult.Item {
			log.Printf("got item:%v \n", v)
		}
	}
}

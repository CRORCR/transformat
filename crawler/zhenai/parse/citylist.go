package parse

import (
	"regexp"
	"transformat/crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	//<a href="http://www.zhenai.com/zhenghun/zigong"
	//class="">自贡</a>
	compile := regexp.MustCompile(cityListRe)
	all := compile.FindAllSubmatch(contents, -1)
	//对于每一个match(city和url),每一个match中的url生成一个新的request
	result := engine.ParseResult{} //城市名称作为result
	for _, v := range all {
		//存储城市列表到item
		result.Item = append(result.Item, string(v[2]))
		//存储所有请求到request
		result.Request = append(result.Request, engine.Request{
			Url:       string(v[1]),
			ParseFunc: engine.NewParse,
		})
		//fmt.Printf("City:%s	URL:%s \n", v[2], v[1])
	}
	return result
}

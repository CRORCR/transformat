package parse

import (
	"fmt"
	"io/ioutil"
	"testing"
	"transformat/crawler/fetcher"
)

//把网页数据备份到本地,就不会依赖网络了
func backUpText(url string) {
	body, err := fetcher.Fetcher(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", body)
}
func TestParseCityList(t *testing.T) {
	//backUpText("http://www.zhenai.com/zhenghun")
	body, err := ioutil.ReadFile("cityList_data.txt")
	if err != nil {
		panic(err)
	}
	parseResult := ParseCityList(body)
	//1.我现在知道结果是470行,先测试是否正确
	const resultSize = 470
	if len(parseResult.Request) != resultSize {
		t.Errorf("result should have:%d,but had:%d", resultSize, len(parseResult.Request))
	}
	//2.检验前三个城市和url是否正确
	expectUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectCitys := []string{"阿坝", "阿克苏", "阿拉善盟"}
	for i, url := range expectUrls {
		if parseResult.Request[i].Url != url {
			t.Errorf("expect url:#%d:%s;but was:s", i, url, parseResult.Request[i].Url)
		}
	}
	for i, city := range expectCitys {
		//知道这是string类型,就可以强转
		if parseResult.Item[i].(string) != city {
			t.Errorf("expect url:#%d:%s;but was:s", i, city, parseResult.Item[i].(string))
		}
	}
	t.Log("success")
}

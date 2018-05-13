package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

//输入是url 输出是utf8编码的文本
func Fetcher(url string) ([]byte, error) {
	//1.获得页面
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//2.结束需要close body
	defer resp.Body.Close()

	//3.获得页面内容
	//判断状态码是ok
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code:%d", resp.StatusCode)
	}
	e := determineEncoding(resp.Body)
	//4.1 gbk读取数据 Decoder解码
	reader := transform.NewReader(resp.Body, e.NewDecoder())
	//通用性比较差 硬编码
	//reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())

	//4.读取所有数据,并输出
	return ioutil.ReadAll(reader)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	//determine读取1024字节去测试编码,如果直接使用r读取,那这1024读完就丢了,
	// 造成数据不完整,所以使用bufio peek读取,返回是新的数组
	bytes, e := bufio.NewReader(r).Peek(1024)
	//如果不知道什么编码,就返回utf8编码
	if e != nil {
		log.Printf("fetcher err:%v \n", e)
		return unicode.UTF8
	}
	encoding, _, _ := charset.DetermineEncoding(bytes, "")
	return encoding
}

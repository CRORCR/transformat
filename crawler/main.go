package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	//1.获得页面
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	//2.结束需要close body
	defer resp.Body.Close()

	//3.获得页面内容
	//判断状态码是ok
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:status code", resp.StatusCode)
		return
	}
	e := determineEncoding(resp.Body)
	//4.1 gbk读取数据 Decoder解码
	reader := transform.NewReader(resp.Body, e.NewDecoder())

	//4.读取所有数据,并输出
	all, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	//determine读取1024字节去测试编码,如果直接使用r读取,那这1024读完就丢了,造成数据不完整,所以使用bufio peek读取,返回是新的数组
	bytes, e := bufio.NewReader(r).Peek(1024)
	if e != nil {
		panic(e)
	}
	encoding, _, _ := charset.DetermineEncoding(bytes, "")
	return encoding
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"transformat/shortUrl/model"
)

func main() {
	shortUrl:=getShortUrl()
	fmt.Printf("short url:%s \n","http://localhost:8080/"+shortUrl)

	originUrl:=getOriginUrl(shortUrl)
	fmt.Printf("origin url:%s \n",originUrl)

}

func getShortUrl()(shortUrl string){
	var buff bytes.Buffer
	var long2Short model.Long2ShortUrlRequest
	long2Short.OriginUrl="https://www.douban.com/note/659541694/"
	data,err:=json.Marshal(long2Short)
	if err!=nil{
		fmt.Printf("marshal failed,err:%v \n",err)
		return
	}
	buff.WriteString(string(data))
	response,err:=http.Post("http://localhost:8080/trans/long2short","application/json",&buff)
	if err!=nil{
		fmt.Printf("post is failed,err:",err)
		return
	}
	result,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		fmt.Printf("post is failed,err:",err)
		return
	}

	var reponse model.Long2ShortUrlReponse
	json.Unmarshal(result,&reponse)

	shortUrl=reponse.ShortUrl
	fmt.Printf("hello world,%v",shortUrl)
	return
}

func getOriginUrl(shortUrl string)(originUrl string){
	var buff bytes.Buffer
	var short model.Short2LongRequest
	short.ShortUrl=shortUrl

	data,err:=json.Marshal(short)
	if err!=nil{
		fmt.Println("marshal is failed",err)
		return
	}
	buff.WriteString(string(data))
	response,err:=http.Post("http://localhost:8080/trans/short2long","application/json",&buff)
	if err!=nil{
		fmt.Println("短转长回报失败",err)
		return
	}

	var origin model.Short2LongUrlReponse

	dataByte,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		fmt.Println("读取body失败",err)
		return
	}
	json.Unmarshal(dataByte,&origin)
	originUrl=origin.OriginUrl
	return
}
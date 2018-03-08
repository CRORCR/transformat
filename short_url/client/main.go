package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"transformat/short_url/model"
)

func main() {
	shortUrl:=getShortUrl()
	fmt.Printf("short url:%s \n","http://localhost:8080/"+shortUrl)

	originUrl:=getOriginUrl(shortUrl)
	fmt.Printf("origin url:%s \n",originUrl)
}

func getShortUrl()(shortUrl string){
	var buffer bytes.Buffer
	var req model.Long2ShortRequest
	req.OriginUrl="https://www.douban.com/note/659541693/"

	data,_:=json.Marshal(req)
	buffer.WriteString(string(data))

	response,err:=http.Post("http://localhost:8080/trans/long2short","application/json",&buffer)
	if err!=nil{
		fmt.Printf("post is failed,err:",err)
		return
	}
	result,_:=ioutil.ReadAll(response.Body)
	fmt.Println(string(result))

	//短找长地址
	var reponse model.Long2ShortResponse
	json.Unmarshal(result,&reponse)
	shortUrl=reponse.ShortUrl
	return
}

func getOriginUrl(short string)(originUrl string){
	var buffer bytes.Buffer
	var req model.Short2LongRequest
	req.ShortUrl=short

	data,_:=json.Marshal(req)
	buffer.WriteString(string(data))

	response,err:=http.Post("http://localhost:8080/trans/short2long","application/json",&buffer)
	if err!=nil{
		fmt.Printf("post is failed,err:",err)
		return
	}
	result,_:=ioutil.ReadAll(response.Body)
	fmt.Println(string(result))

	//短找长地址
	var reponse model.Short2LongResponse
	json.Unmarshal(result,&reponse)
	originUrl=reponse.OriginUrl
	return
}

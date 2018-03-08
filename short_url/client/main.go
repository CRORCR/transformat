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
	fmt.Printf("short url:%s \n","http://localhost:8080/"+reponse.ShortUrl)
}

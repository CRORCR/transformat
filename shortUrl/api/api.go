package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"transformat/shortUrl/logic"
	"transformat/shortUrl/model"
)

const(
	ErrSuccess=0
	ErrInvalidParameter=1001
	ErrServerBusy= 1002
)
func main() {
	err:=logic.InitDb("root:root@tcp(127.0.0.1:3306)/short_url?parseTime=true")
	if err!=nil{
		fmt.Println("连接数据库有误:",err)
	}
	http.HandleFunc("/trans/long2short",Long2Short)
	http.HandleFunc("/trans/short2long",Short2Long)
	http.ListenAndServe(":8080",nil)
}

//长url转短url
//1.转成长url对象
//2.转成md5加密,去数据库查
//3.查到就返回 短url
//4.没有查到就存入数据库,再返回
func Long2Short(w http.ResponseWriter,r *http.Request){
	//1.读取数据
	data,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		reponseErr(w,1001)
		fmt.Println("读取文件有误,err:",err)
		return
	}
	longReq:=&model.Long2ShortUrlRequest{}
	err=json.Unmarshal(data,longReq)
	if err!=nil{
		fmt.Println("反序列化失败,err:",err)
		reponseErr(w,1002)
		return
	}

	reponse,err:=logic.Long2Short(longReq)
	fmt.Println("reponse值",reponse)
	if err!=nil{
		fmt.Println("回报失败,err:",err)
		return
	}
	reponseSucc(w,reponse)
}

func reponseErr(w http.ResponseWriter ,code int){
	var head model.Header
	head.Code=code
	head.Message=getMessage(code)
	data,err:=json.Marshal(head)
	if err!=nil{
		s:="{\"code\":500,\"message\":\"server is failed\"}"
		w.Write([]byte(s))
		return
	}
	w.Write(data)
}

func reponseSucc(w http.ResponseWriter,data interface{}){

	dataByte,err:=json.Marshal(data)
	if err!=nil{
		s:="{\"code\":500,\"message\":\"server is failed\"}"
		w.Write([]byte(s))
		return
	}
	w.Write(dataByte)
}

func getMessage(code int)(message string){

	switch code {
	case 1001:
		message="ErrInvalidParameter"
	case 1002:
		message="ErrServerBusy"
	default:
		message="unknow error"
	}
	return
}

func Short2Long(w http.ResponseWriter,r *http.Request){
	data,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("短url读取body失败,err:",err)
		return
	}
	var shortReq *model.Short2LongRequest
	json.Unmarshal(data,&shortReq)

	reponse,err:=logic.Short2Long(shortReq)
	fmt.Println("reponse值",reponse)
	if err!=nil{
		fmt.Println("回报失败,err:",err)
		return
	}
	reponseSucc(w,reponse)
}
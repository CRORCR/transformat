package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"transformat/short_url/logic"
	"transformat/short_url/model"
)
//定义错误常量
const(
	ErrSuccess=0
	ErrInvalidParameter=1001
	ErrServerBusy= 1002
)

func main() {
	err:=logic.InitDb("root:root@tcp(127.0.0.1:3306)/short_url?parseTime=true")
	if err!=nil{
		fmt.Println("初始化失败",err)
		return
	}
	http.HandleFunc("/trans/long2short",long2short)
	http.ListenAndServe(":8080",nil)
}

func long2short(w http.ResponseWriter,r *http.Request){
	data,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println("read request is failed")
		reposneErr(w,1001)
		return
	}
	var req model.Long2ShortRequest
	err=json.Unmarshal(data,&req)
	if err!=nil{
		fmt.Println("unmarshal is failed",err)
		reposneErr(w,1002)
		return
	}
	resp,err:=logic.Long2Short(&req)
	if err!=nil{
		fmt.Println("reposne is failed",err)
		reposneErr(w,1002)
		return
	}
	reposneSuccess(w,resp)
}

func reposneErr(w http.ResponseWriter,code int){
	var reponse model.Header
	reponse.Code=code
	reponse.Message=getMessage(code)
	data,err:=json.Marshal(reponse)
	if err!=nil{
		w.Write([]byte("{\"code\":500, \"message\": \"server busy\"}"))
		return
	}
	w.Write(data)
}
func reposneSuccess(w http.ResponseWriter,data interface{}){
	dataByte,err:=json.Marshal(data)
	if err!=nil{
		s:="{\"code\":500, \"message\": \"server busy\"}"
		w.Write([]byte(s))
		return
	}
	w.Write(dataByte)
}

func getMessage(code int)(mes string){
	switch code {
	case ErrSuccess:
		mes= "success"
	case ErrInvalidParameter:
		mes= "invalid parameter"
	case ErrServerBusy:
		mes= "server busying"
	default:
		mes= "unkown error"
	}
	return
}
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const form = `<html>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"> 
<body><form action="#" method="post" name="bar">
<div> 姓名：<input type="text" name="username"/></div>
<div> 密码：<input type="text" name="password"/></div>
<input type="submit" value="登录"/>
</form></html></body>`

func main() {
	http.HandleFunc("/test1", SimpleServer)
	http.HandleFunc("/test2", FormServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Printf("listen is failed err;%v \n", err)
		return
	}
}

func SimpleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>hello world</h1>")
}

func FormServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") //设置相应头,可以自定义
	//w.Header().Set("Server","GO 1.9.2")
	//w.Header().Set("shabi","hhh")//自定义
	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		r.ParseForm()
		//多个name一样,可以使用角标获取
		//io.WriteString(w, r.Form["in"][0])
		//io.WriteString(w, r.Form["in"][1])
		io.WriteString(w, r.FormValue("username"))
		io.WriteString(w, "\n")
		io.WriteString(w, r.FormValue("password"))
	}
}

//捕获异常
func logPanics(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		handle(writer, request)
	}
}

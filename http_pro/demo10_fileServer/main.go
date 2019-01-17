package main

import (
	"fmt"
	//"os"
	"html/template"
	"net/http"
)

var (
	gtemp *template.Template
)

func init() {
	t, err := template.ParseFiles("./day10/demo10_fileServer/views/index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	gtemp = t
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	gtemp.Execute(w, nil)
}

func main() {
	//处理请求
	http.HandleFunc("/index", handleIndex)
	//加载静态资源
	//http.Handle("/static/", http.FileServer(http.Dir("./static/")))
	http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("./day10/demo10_fileServer/media/"))))
	http.ListenAndServe(":8088", nil)
}

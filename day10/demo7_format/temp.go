package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("K:/workspace/src/transformat/day10/demo7_format/index.html")
	if err != nil {
		fmt.Printf("parse file failed ,err :%v \n", err)
		return
	}
	p := Person{"lcq", 8}
	if err := t.Execute(w, p); err != nil {
		fmt.Printf("there was an error:%v \n", err.Error())
		return
	}
}

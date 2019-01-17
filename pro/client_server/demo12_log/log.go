package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/helloworld", logHeand(helloworld))
	http.ListenAndServe(":8080", nil)
}

func helloworld(w http.ResponseWriter, r *http.Request) {
	string := "hello world,你好,世界!!!"
	w.Write([]byte(string))
}

func logHeand(function func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UnixNano()
		helloworld(w, r)
		end := time.Now().UnixNano()
		count := (end - start) / 1000
		log.Printf("url:%s,count:%d,agent:%s", r.RequestURI, count, r.UserAgent())
	}
}

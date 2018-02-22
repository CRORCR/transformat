package main

import (
	"fmt"
	"io"
	"net/http"
)

const form = `<html><body><form action="#" method="post" name="bar">
<input type="text" name="in"/>
<input type="text" name="in"/>
<input type="submit" value="Submit"/>
</form><html></body>`

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
		io.WriteString(w, r.Form["in"][0])
		io.WriteString(w, r.Form["in"][1])
		io.WriteString(w, "\n")
		io.WriteString(w, r.FormValue("in"))
	}
}

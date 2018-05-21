package main

import (
	"fmt"
	"net/http"
)

//服务端
func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("listen is failed")
	}

}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
	fmt.Fprintf(w, "hello world")
}

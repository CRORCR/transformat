package main

import (
	"fmt"
	"transformat/basic/demo5_inter/mock"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("www.imooc.com")
}

func main() {
	r := mock.Retriener{"hello"}
	fmt.Printf("%T", r)
	fmt.Println(download(r))
}

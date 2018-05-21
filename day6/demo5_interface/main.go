package main

import (
	"fmt"
)

//空接口是所有类型的父类
func main() {
	var a interface{}
	var b int = 1000

	a = b
	fmt.Println(a)

	var c string = "hello"
	a = c
	fmt.Println(a)
}

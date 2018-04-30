package main

import (
	"fmt"
	"strings"
)

//闭包
//
//输出:1 101 1101
func main() {
	//demo()
	demo2()
}

func demo(){
	f := Adder()
	fmt.Println(f(1))
	fmt.Println(f(100))
	fmt.Println(f(1000))
}
func Adder() func(int) int {
	var x int
	f := func(d int) int {
		x += d
		return x
	}
	return f
}

func demo2(){
	f := makeSuffix(".jpg")
	fmt.Println(f("test"))
	fmt.Println(f("pic"))
}
//是否以什么后缀结尾,不是就添加
func makeSuffix(suffix string) func(string) string {
	f := func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
	return f
}

package main

import (
	"fmt"
)

//iota是从0开始计数的计数器
//没有iota就使用上行变量
//再次遇到iota还是从头计数
//遇到const才会重新从0计数
const (
	a  = iota
	b  = iota
	c  = iota
	d1 = 2
	d2
	d3
	e1 = iota
	e2 = iota
	e3 = iota
)

func main() {
	fmt.Println(a) //0
	fmt.Println(b) //1
	fmt.Println(c) //2

	fmt.Println(d1) //2
	fmt.Println(d2) //2
	fmt.Println(d3) //2

	fmt.Println(e1) //6
	fmt.Println(e2) //7
	fmt.Println(e3) //8
}


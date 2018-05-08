package main

import "fmt"

//type自定义类型
type Int int

func main() {
	var a Int
	var b int
	//Int 和 int虽然底层存储是一样的,但是不能直接赋值,必须要强制转换一下
	a = Int(b)
	fmt.Println(a)
}
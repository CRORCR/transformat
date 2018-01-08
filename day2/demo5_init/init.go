package main

import "fmt"

//init函数是在main函数之前,执行,如果其他文件中也有int函数,一并执行
//在所有的init函数中,main函数所在包下的init最先执行,然后执行其他文件
//中的init函数

//go build  然后执行demo5_init文件
//init 函数先执行
//other init 函数
//mian函数后执行

func init() {
	fmt.Println("init 函数先执行")
}
func main() {
	fmt.Println("mian函数后执行")
}

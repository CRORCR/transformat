package main

import "fmt"

func main() {
	skiphand()
	//nullhand()
}

//如果是0,继续往下执行,不管下一条是什么都会执行(傻逼语法,不要用)
//hello
//1
func skiphand(){
	var i = 0
	switch i {
	case 0:
		fmt.Println("hello")
		fallthrough
	case 1:
		fmt.Println("2")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}
}

//如果是0 就跳过,什么也不做
//case 2,3,4,5 知识点,满足任何一个都会执行
func nullhand(){
	var i = 0
	switch i {
	case 0:
	case 1:
		fmt.Println("1")
	case 2,3,4,5:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}
}
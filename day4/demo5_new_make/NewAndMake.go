package main

import "fmt"

func main() {
	test()
}

/*
输出:
	&[]
	&[100 0 0 0 0]
	[0 0 0 0 0 0 0 0 0 0]
*/
func test() {
//new返回的是指针
	s1 := new([]int)
	fmt.Println(s1)
	//new返回是指针,没有初始化,获取值,就肯定panic
	//(*s1)[0]=100
	*s1 = make([]int, 5)
	(*s1)[0] = 100
	fmt.Println(s1)

	//make返回是引用
	s2 := make([]int, 10)
	fmt.Println(s2)
	return
}
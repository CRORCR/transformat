package main

import "fmt"

func main() {
	//demo()
	demo2()
}
func demo() {
	a := 100
	var p *int
	p = &a
	*p = 200
	fmt.Println(*p) //a值
	fmt.Println(p)  //a地址
	fmt.Println(a)  //a值
	fmt.Println(&a) //a地址
}

func demo2() {
	//指针不能直接使用
	//var p *int
	//*p = 100

	//1.第一种方式,可以把变量的地址给指针变量
	var p1 *int
	str := 123
	p1 = &str

	//2.使用new创建一个指针才能使用
	var p2 = new(int)
	*p2 = 100

	fmt.Println(p1, p2)
}

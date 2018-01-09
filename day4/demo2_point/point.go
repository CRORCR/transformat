package main

import "fmt"

func main() {
	//test1()
	//test2()
	test3()

}

func test1() {
	a := 10
	var p *int
	p = &a
	fmt.Printf("%p \n", &p) //输出变量p的地址
	fmt.Printf("%p", p)     //输出p存的指针地址

	fmt.Printf("%p \n", &a)
	fmt.Printf("%p", a)
}

func modify(a *int) {
	*a = 100
}

func test2() {
	var b int = 1
	var p *int
	p = &b
	modify(p)      //和modify(&b) 效果一样
	fmt.Println(b) //100
}

func test3() {

	///p 默认初始化nil
	var p *int
	var b int
	p = &b
	*p = 200 //b = 200

	if p == &b {
		fmt.Printf("equal\n")
	}

	fmt.Printf("%p %p %p\n", p, &b, &p)

	//使用new创建,默认返回地址值
	p = new(int)
	*p = 1000
	fmt.Printf("%d\n", *p)
	fmt.Printf("%p %p %p\n", p, &b, &p)

	if p == &b {
		fmt.Printf("equal")
	}

	//指针类型的变量初始化：1. 使用其他变量地址给它赋值。 2. 使用new分配
}

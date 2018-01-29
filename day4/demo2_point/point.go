package main

import "fmt"

func main() {
	demo1()
	//test2()
	//test3()
	//test4()
	//test5()
}
func demo1() {
	a := 100
	var p *int
	fmt.Printf("%p \n", &p) //地址
	fmt.Printf("%p \n", p)  //p地址,指针地址0x0
	fmt.Printf("%d \n", a)  //a值
	fmt.Printf("%p \n", &a) //a地址

	p = &a
	fmt.Printf("%d\n", *p) //100
	*p = 200
	fmt.Println(a, *p) //200 200
}

func test2() {

	//p 默认初始化nil
	var p *int
	var b int
	p = &b
	*p = 200 //b = 200

	if p == &b {
		fmt.Printf("equal\n") //equals
	}

	fmt.Printf("%p %p %p\n", p, &b, &p) //b地址  b地址 p地址

	p = new(int)
	*p = 1000
	fmt.Printf("%d\n", *p)              //1000
	fmt.Printf("%p %p %p\n", p, &b, &p) //new的地址 // b地址 p地址

	if p == &b {
		fmt.Printf("equal")
	}
}

func test3() {
	var p *string = new(string)
	*p = "hello"
	fmt.Printf("%p \n", p)  //0xc04203e1b0
	fmt.Printf("%p \n", &p) //0xc042068018

	var s string = "hello"
	p = &s
	fmt.Printf("%s \n", *p) //hello
	fmt.Printf("%p \n", p)  //0xc04203e1b0
	fmt.Printf("%p \n", &p) //0xc042068018
}

func test4() {
	var a []int = make([]int, 10)
	a[0] = 100
	fmt.Println(a) //[100 0 0 0 0 0 0 0 0 0]

	var p *[]int = new([]int)
	(*p) = make([]int, 10)
	(*p)[0] = 100
	fmt.Println(p) //&[100 0 0 0 0 0 0 0 0 0]

	p = &a
	(*p)[0] = 1000
	fmt.Println(a) //[1000 0 0 0 0 0 0 0 0 0]
}

//切片是指向数组的,切片值改变,对应数组值也会改变
func test5() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	b := arr[:]
	b[0] = 100
	fmt.Println(arr, b)
}

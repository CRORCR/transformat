package main

import "fmt"

func main() {
	test1()
}

func test1() {
	var a int = 10
	var b *int
	fmt.Printf("%p\n", &b) //0xc042004028
	fmt.Printf("%p\n", &a) //0xc0420080a8
	fmt.Printf("%p\n", b)  //0x0
	b = &a
	fmt.Printf("%d\n", *b)
	*b = 100
	fmt.Printf("a=%d\n", a)
}

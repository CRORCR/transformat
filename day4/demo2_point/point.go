package main

import "fmt"

func main() {
	demo()
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
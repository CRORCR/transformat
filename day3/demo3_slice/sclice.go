package main

import "fmt"

func main() {
	testSlice()
}

func testSlice() {
	var a = [10]int{1, 2, 3, 4}

	b := a[1:5]
	fmt.Printf("%p\n", b)     //2 3 4 0
	fmt.Printf("%p\n", &a[1]) //切片指向数组第一个元素地址
}

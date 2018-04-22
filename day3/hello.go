package main

import "fmt"

func main() {
	//test()
}

//1.数组是值copy,改变后,原数组不变
//2.切片改变数组会随着改变
//3.切片长度超过后,会重新分配地址,这时候改变切片,原数组不变
func test() {
	var a []int = make([]int, 10)
	var b [10]int = [10]int{1, 2, 3, 8: 100}
	fmt.Println(b) //b数组 1 2 3 0 0 0 0 0 100 0 0

	var c [10]int = b //值传递 copy
	fmt.Println(c)    //1 2 3 0 0 0 0 0 100 0 0
	c[0] = 1000       //c改变b不变
	fmt.Println(b)    //1 2 3 0 0 0 0 0 100 0 0

	a = b[:]
	a[0] = 1000
	fmt.Println(b) //1000 2 3 0 0 0 0 0 100 0 0

	a = append(a, 10, 30, 40)
	fmt.Println(a) //1000 2 3 0 0 0 0 0 100 0 0 10, 30, 40
	a[0] = 2000

	fmt.Println(b) //1000 2 3 0 0 0 0 0 100 0 0
}

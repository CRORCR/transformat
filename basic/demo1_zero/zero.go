package main

import (
	"fmt"
	"strconv"
)

func main() {
	//zero()
	fmt.Println(converToBin(5))
	//demo()
}
func zero() {
	var a int
	var b string
	//fmt.Println(a, b) //0 后面空格不会输出
	fmt.Printf("%d %q", a, b) //0 ""   %q可以输出字符串引号
}

//10进制转为2进制
func converToBin(n int) (result string) {
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return
}

func demo() {
	fmt.Println(0 | 0)
}

package main

import (
	"fmt"
)

//阶乘
func main() {
	num := jiecheng(3)
	fmt.Println(num)
}

func jiecheng(i int) int {
	if i == 1 {
		return 1
	} else {
		return i * jiecheng(i-1)
	}
}

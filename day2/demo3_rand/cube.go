package main

import (
	"fmt"
	"math"
)

var numList []int = make([]int, 10)

//水仙花数
//计算存在一个三位数,每一位的三次方,加起来等于这个数
func main() {
	numList = hello()
	for _, v := range numList {
		fmt.Println(v)
	}
}

func hello() (numList []int) {
	for i := 100; i < 1000; i++ {
		//个位
		ge := float64(i % 100 % 10)
		//十位
		shi := float64(i / 10 % 10) // %取小数点后面值
		//百位
		bai := float64(i / 100) //  / 取小数点前面值

		if math.Pow(ge, 3)+math.Pow(shi, 3)+math.Pow(bai, 3) == float64(i) {
			numList = append(numList, i)
		}
	}
	return
}

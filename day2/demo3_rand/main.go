package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	//计算存在一个三位数,每一位的三次方,加起来等于这个数
	//for i := 100; i < 999; i++ {
	//	if isNarcissistic(i) {
	//		fmt.Println(i)
	//	}
	//}
	demo()
}

func isNarcissistic(n int) bool {
	//得出百位 十位 个位  数字
	i, j, k := float64(n/100), float64(n/10%10), float64(n%10)
	//求和,等于这个三位数
	if math.Pow(i, 3)+math.Pow(j, 3)+math.Pow(k, 3) == float64(n) {
		return true
	}
	return false
}

func demo() {
	//随机数  intn求在某范围内的数字
	fmt.Println(rand.Intn(100)) //81
	//随机数 求在float64范围内的数字
	fmt.Println(rand.Float64()) //0.9405090880450124
	fmt.Println(rand.Float32()) //0.6645601
}

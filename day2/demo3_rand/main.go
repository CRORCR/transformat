package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	demo()
}

func demo() {
	//随机数  intn求在某范围内的数字
	fmt.Println(rand.Intn(100)) //81
	//随机数 求在float64范围内的数字
	fmt.Println(rand.Float64()) //0.9405090880450124
	fmt.Println(rand.Float32()) //0.6645601
}

package main

import "fmt"

func main() {
	//rangeDemo()
	//printDemo()
	//fblq1()
	fblq2()
}

//range会值拷贝,如果修改数组,对原始数据没有任何影响,应该通过角标操作数据
//不应该直接操作value
func rangeDemo() {
	//arr := [4]int{1, 2, 3, 4}
	//for _, v := range arr {
	//	v = v + 1
	//}
	//fmt.Println(a)

	arr := [4]int{1, 2, 3, 4}
	for i, _ := range arr {
		arr[i] = arr[i] + 1
	}
	fmt.Println(arr)
}

func printDemo() {
	fmt.Println(3 / 2)      //1
	fmt.Println(3.0 / 2)    //1.5
	fmt.Printf("%f\n", 3/2) //使用%f不管用
	//%f只能用于确定是浮点类型
	var ff float32
	fmt.Printf("%f\n", ff) //0.000000
}

var num = 5
var sum = 0

//斐波那契数列_递归实现
func fblq1() {
	for i := 1; i <= num; i++ {
		sum += fblq11(i)
	}
	fmt.Println(sum)
}

func fblq11(i int) int {
	fmt.Println(i, sum)
	if i <= 1 {
		return 1
	}
	if i == 2 {
		return 1
	}
	return fblq11(i-1) + fblq11(i-2)
}

//斐波那契数列_切片实现
func fblq2() {
	sum2 := fbnq22(5) //[1 1 2 3 5]
	fmt.Println(sum2)
}

func fbnq22(tt int) []int {
	arr := make([]int, tt)
	arr[0], arr[1] = 1, 1
	for i := 2; i < tt; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr
}

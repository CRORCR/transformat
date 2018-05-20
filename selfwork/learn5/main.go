package main

import "fmt"

func main() {
	rangeDemo()
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

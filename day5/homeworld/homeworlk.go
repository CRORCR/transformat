package main

import "fmt"

//
func main() {
	//maopao()
	bijiao()
}

func maopao() {
	var arr = []int{10, 5, 3, 8, 2, 9}
	for i := len(arr); i >= 0; i-- {
		for j := 1; j < i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

//选择
func bijiao() {
	var arr = []int{10, 5, 3, 8, 2, 9}
	for i := 0; i < len(arr)-1; i++ {
		for j := 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
			fmt.Println(arr)
		}
		fmt.Println(arr)
	}
	//fmt.Println(arr)
}

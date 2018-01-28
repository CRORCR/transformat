package main

import "fmt"

//反转
func main() {
	reverStr()
	//reverInt()
	//reverSlice()
}

//string反转  先转为数组,再转回字符串
func reverStr() {
	str := "hello world"
	b := []byte(str)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	str = string(b)
	fmt.Println(str)
}

//切片反转
func reverSlice() {
	b := make([]int, 4)
	b[0] = 1
	b[1] = 2
	b[2] = 3
	b[3] = 4
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	fmt.Println(b)
}

//int反转
func reverInt() {
	a, b := 1, 2
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}

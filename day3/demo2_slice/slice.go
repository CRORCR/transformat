package main

import (
	"fmt"
)

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	test6()
}

//切片的值改变,会改变数组的值
func test1() {
	a := [3]int{1, 2, 3}
	b := a[:]
	b[0] = 100
	fmt.Println(a) //[100 2 3]
	fmt.Println(b) //[100 2 3]
}
func test2() {
	a := make([]int, 5, 10)
	a[4] = 100
	b := a[2:3]
	//a=[]int{0, 0, 0, 0, 100}, len(a) = 5, cap(a)=10
	fmt.Printf("a=%#v, len(a) = %d, cap(a)=%d\n", a, len(a), cap(a))
	//b=[]int{0}, len(b) = 1, cap(b)=8
	fmt.Printf("b=%#v, len(b) = %d, cap(b)=%d\n", b, len(b), cap(b))
}

//append 添加切片  不能添加数组
func test3() {
	a := make([]int, 5)
	b := []int{1, 2}
	a = append(a, b...)
	fmt.Println(a)
	fmt.Println(b)
}

func test4() {
	//定义字符串
	str := "hello world"
	//转换成数组切片
	b := []byte(str)
	b[0] = 'Q'
	string := string(b)
	fmt.Println(string) //Qello world
}

//字符串翻转
func test5() {
	str := "hello world"
	b := []byte(str)

	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	str1 := string(b)
	fmt.Println(str1)
}

func test6() {
	str := "hello world我们爱中国"
	b := []rune(str)

	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	str1 := string(b)
	//国中爱们我dlrow olleh
	fmt.Println(str1)
	//len(str)=26, len(rune)=16
	fmt.Printf("len(str)=%d, len(rune)=%d\n", len(str), len(b))
}

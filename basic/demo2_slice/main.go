package main

import "fmt"

var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}

func main() {
	//sliceDemo()
	//appendDemo()
	deleteDemo()
}

func deleteDemo() {
	//如何删除第三个角标的元素
	fmt.Println(arr)                 //0 1 2 3 4 5 6 7
	s := append(arr[:3], arr[4:]...) //0 1 2 4 5 6 7
	fmt.Println(s)
}

//往切片添加元素的时候,如果长度超过cap,系统会重新分配更大的底层数组
//原来的数组怎么办?
//go语言有垃圾回收机制,如果有程序用就不做任何操作,如果没有使用,系统自动回收
func appendDemo() {
	s1 := arr[2:6]       //2, 3, 4, 5
	s2 := s1[3:5]        //5 6
	s3 := append(s2, 10) //5 6 10
	s4 := append(s3, 11) //5 6 10 11
	s5 := append(s4, 12) //5 6 10 11 12
	fmt.Println(s1, s2, s3, s4, s5)
	fmt.Println(arr) //0 1 2 3 4 5 6 10
}

//为什么s2没有报错,从s1取出单个值会报错呢?
//因为切片是指向数组的,s1是指向了数组角标为2的元素,它知道后面还是有元素的,但是不能直接取出来
func sliceDemo() {
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println(s1)    //[2 3 4 5]
	fmt.Println(s2)    //[5 6]
	fmt.Println(s1[4]) //index out of range
}

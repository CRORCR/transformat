package main

import (
	"fmt"
)

var a int

//可变参数
func main() {
	//test1()
	//test3()
}
func test1() {
	fmt.Println(demo())           //0
	fmt.Println(demo(1, 2, 3, 4)) //10

}

//多个参数,一般是放在参数列表最后面,多参数包括参数个数为0
func demo(args ...int) (sum int) {
	for _, v := range args {
		sum += v
	}
	return
}

//defer
//defer是在最后执行,return之前执行
//多个defer执行是按照先进后出的顺序执行
func test3() {
	a := 100
	fmt.Println(a)                    //1
	defer fmt.Println("第一个defer:", a) //4

	b := 200
	fmt.Println(b)                    //2
	defer fmt.Println("第二个defer:", b) //3
	if b == 200 {
		return
	}
}

package main

import (
	"fmt"
	"math/rand"
)
//注意:
//数组在当做值传递的时候,长度不一样,默认是不一样的类型,不能传递

//地址值使用%p 取出数组值 使用&赋值输出
func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//genRand()
	genRandStr()
}

//定义一个数组,输出每个值的地址值
func test1() {
	var arr [10]int8
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%p \n", &arr[i]) //输出10个地址值
	}
	for index, _ := range arr { //遍历输出值
		fmt.Printf("a[%d]=%d\n", index, arr[index])
	}
}

//定义一个数组a,赋值给数组b,然后修改b的值,发现a没有变化,说明数组是值传递
func test2() {
	a := [3]int{1, 2, 3}
	b := a
	b[0] = 100
	fmt.Println(a) //[1 2 3]
	fmt.Println(b) //[100 2 3]
}

//定义数组的几种形式
//var a [3] int = [3]int{1,2,3,}
//var b = [...]int{4,5,6}
// c := [...]int{1,2,3}   最常用
//如果给定了长度,元素不够,补充默认值0
// 也可以给定角标赋值
//var d [5]string = [5]string{1:"abc", 4:"efg"}
func test3() {
	var a [3]int = [3]int{1, 2, 3}
	fmt.Printf("%v\n", a) //[1 2 3]

	var b = [...]int{4, 5, 6}
	fmt.Printf("%v\n", b) //[4 5 6]

	var c = [4]int{1, 2, 3}
	fmt.Printf("%v\n", c) //[1 2 3 0]

	var d [3]string = [3]string{1: "aaa", 2: "bbb"}
	fmt.Printf("%#v\n", d) //[3]string{"", "aaa", "bbb"}
}

//多维数组  赋值和打印
func test4() {
	var a [4][2]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] = (i + 1) * (j + 1)
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}
}

func genRand() {
	var a [10]int
	for i := 0; i < len(a); i++ {
		//赋值
		a[i] = rand.Int() //int范围内的值
		//a[i] = rand.Intn(100)//100以内的值
	}
	//遍历输出随机数
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d\n", a[i])
	}
}

func genRandStr() {
	var arr = [10]string{}
	count := "helloworld"
	for i := 0; i < len(count); i++ {
		index := rand.Intn(len(count))
		//str := fmt.Sprintf("%s", count[index])
		str := fmt.Sprintf("%s", string(count[index]))
		if i == 10 {
			break
		}
		arr[i] = str

	}
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Println(len(arr))
}

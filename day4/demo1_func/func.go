package main

import (
	"fmt"
	"sort"
)

var a int

func main() {
	//test1()

	//test3()

	testMap()
}
func test1() {
	fmt.Println(demo())           //0
	fmt.Println(demo(1))          //1
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
func testMap() {
	var a map[string]int
	a = make(map[string]int, 100)
	a["abc"] = 0
	a["hello"] = 1200
	a["cello"] = 1200
	fmt.Println(a)

	var keys []string

	for k, v := range a {
		fmt.Printf("a[%s] = %d\n", k, v)
		keys = append(keys, k)
	}

	fmt.Println("\n")
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("a[%s]=%d\n", k, a[k])
	}

	//第一种写法，返回val用来接收key "abc"的值，exist对应key "abc"是否在a中存在。
	val, ok := a["abc"]
	fmt.Printf("val=%d ok = %t\n", val, ok)
	if ok {
		fmt.Printf("val = %d\n", val)
	} else {
		fmt.Println("not found")
	}

	//第二种写法，val直接获取key aaa得值，如果aaa不存在，则val为0。这种写法无法区分aaa是否存在
	val = a["aaa"]
	fmt.Println(val)
}

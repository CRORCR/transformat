package main

import (
	"fmt"
)

func main() {
	//for {
	//	test()
	//	time.Sleep(time.Second)
	//}
}

func test() {
	//defer必须在出错代码之前,否则没用
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic:%v \n", err)
		}
	}()
	var p *int
	set(p)

}
func set(p *int) {
	*p = 123
}

//定义一个空的指针 初始值是nil,如果取值*p就会报空指针异常
func demo() {
	var p *int
	fmt.Println(p)
	fmt.Println(*p)
}

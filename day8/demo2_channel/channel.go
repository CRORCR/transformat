package main

import "fmt"

type student struct {
	name string
}

func main() {
	//定义一个chan管道,存储接口
	var stuChan chan interface{}
	//必须初始化
	stuChan = make(chan interface{}, 10)
	//存入结构
	stu := student{name: "stu01"}
	stuChan <- &stu
	//读取管道内容
	var stu01 interface{}
	stu01 = <-stuChan

	//强转成student
	stu02, ok := stu01.(*student)
	if !ok {
		fmt.Println("can not convert")
		return
	}
	fmt.Println(stu02)
}

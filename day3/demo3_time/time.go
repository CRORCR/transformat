package main

import (
	"fmt"
	"time"
)

func main() {
	//test1()
	//test2()
	test3()
}

//年月日输出,格式化输出
func test1() {
	time := time.Now()
	fmt.Printf("%T \n", time) //time类型  time.Time
	year := time.Year()
	month := time.Month()
	day := time.Day()
	//2018-01-09 11:22:52
	fmt.Printf("%04d-%02d-%02d %02d:%02d:%02d", year, month, day, time.Hour(), time.Minute(), time.Second())
	fmt.Printf("timestamp:%d\n", time.Unix()) //时间戳
}

//2006-01-02 03:04:05 固定写法
//时间的format函数参数固定
func test2() {
	now := time.Now()
	str := now.Format("2006-01-02 03:04:05")
	//format result:2018-01-09 11:26:03
	fmt.Printf("format result:%s\n", str)
}

//计算代码运行时间
func test3() {
	start := time.Now().UnixNano()
	/*
		业务代码
	*/
	time.Sleep(10 * time.Millisecond)
	end := time.Now().UnixNano()
	cost := (end - start) / 1000
	fmt.Printf("cost:%dus\n", cost)
}

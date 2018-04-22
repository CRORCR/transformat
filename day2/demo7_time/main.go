package main

import (
	"fmt"
	"time"
)
//时间格式化表示
//now := time.Now()
//fmt.Println(now.Format(“02/1/2006 15:04”))
//fmt.Println(now.Format(“2006/1/02 15:04”))
//fmt.Println(now.Format(“2006/1/02”))

func main() {
	//当前时间
	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05"))//2018/04/22 23:14:11
	//时间戳
	second := now.Unix()
	//2018-01-08 22:46:42.3452308 +0800 CST m=+0.011000601
	fmt.Println(now)
	//1515422802
	fmt.Println(second)
	//性别随机
	if second%2 == 0 {
		fmt.Println("female")
	} else {
		fmt.Println("man")
	}
}
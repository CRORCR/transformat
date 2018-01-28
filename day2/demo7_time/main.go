package main

import (
	"fmt"
	"time"
)

const (
	Man    = 1
	Female = 2
)

func main() {
	//当前时间
	now := time.Now()
	//时间戳
	second := now.Unix()
	//2018-01-08 22:46:42.3452308 +0800 CST m=+0.011000601
	fmt.Println(now)
	//1515422802
	fmt.Println(second)
	//性别随机
	if second%Female == 0 {
		fmt.Println("female")
	} else {
		fmt.Println("man")
	}
}

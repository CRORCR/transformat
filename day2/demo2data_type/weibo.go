package main

import (
	"fmt"
)

const (
	HongMing = 1 << 0 //1
	DaRen    = 1 << 1 //2
	Vip      = 1 << 2 //4
)

type User struct {
	Name string //姓名
	flag uint8  //等级标签
}

//设置等级标签,如果为true就是设置,false就是取消标签

func SetFlag(u User, isSet bool, flag uint8) User {
	if isSet {
		u.flag = u.flag | flag
	} else {
		u.flag = u.flag ^ flag
	}
	return u
}

func is_flag(u User, flag uint8) bool {
	f := u.flag & flag
	return f == 1
}

func main() {
	user := User{"李长全", 0}
	user = SetFlag(user, true, 1)
	fmt.Println(user)
}

package main

import (
	"fmt"
)

const (
	HongMing = 1 << 0
	DaRen    = 1 << 1
	Vip      = 1 << 2
)

type User struct {
	name string
	flag uint8
}

func set_flag(user User, isSet bool, flag uint8) User {
	if isSet == true {
		user.flag = user.flag | flag
	} else {
		user.flag = user.flag ^ flag
	}
	return user
}

func is_set_flag(user User, flag uint8) bool {
	result := user.flag & flag
	return result == flag
}

func main() {
	var user User
	user.name = "test01"
	user.flag = 0

	user = set_flag(user, true, DaRen)
	result := is_set_flag(user, DaRen)
	fmt.Println(user.flag)
	fmt.Println(DaRen)
	fmt.Printf("user is DaRen:%t\n", result)
	result = is_set_flag(user, HongMing)
	fmt.Println(user.flag)
	fmt.Println(HongMing)
	fmt.Printf("user is hongming:%t\n", result)
}

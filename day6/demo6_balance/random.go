package main

import (
	"math/rand"
)

type RandBalance struct {
}

func (r *RandBalance) DoBalance(addrList []string) string {
	l := len(addrList)
	//以长度为基础,取随机数
	index := rand.Intn(l)
	return addrList[index]
}

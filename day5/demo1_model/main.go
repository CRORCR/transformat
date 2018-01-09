package main

import (
	"fmt"
	"oldboy/day5/model"
)

func main() {
	s := model.NewSchool("北京大学", "北京海淀区")
	fmt.Println(s)
	fmt.Printf(s.GetAddr())
	fmt.Printf(s.GetName())
}

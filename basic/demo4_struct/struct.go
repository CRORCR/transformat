package main

import "fmt"

type Student struct {
	Name string
}

func (s Student) returnName() {
	s.Name = "李长全"
}
func (s *Student) returnNamePtr() {
	s.Name = "李长全"
}

func main() {
	stu := Student{}
	stu.returnName()
	fmt.Println(stu.Name)

	stu.returnNamePtr()
	fmt.Println(stu.Name)
}

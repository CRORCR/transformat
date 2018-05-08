package main

import "fmt"

//链表就是结构体字段是其他结构体的指针类型

type People struct {
	Sex string
	s   *Student
}

type Student struct {
	Name string
	Age  int
	p    *People
}

func main() {
	var peo = &People{Sex: "男", s: new(Student)}
	peo.s.Name = "学生:李长全"
	peo.s.Age = 16
	//&{Sex:男 s:0xc0420443a0}
	fmt.Printf("%+v \n", peo)
}

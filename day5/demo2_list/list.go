package main

import "fmt"

type Teacher struct {
	Name string
	Age  int
	Next *Teacher
}

func NewTeacher(name string, age int) *Teacher {
	p := new(Teacher)
	p.Name = name
	p.Age = age
	return p
}

func createList() {
	//第一种赋值方式
	var header *Teacher = &Teacher{}
	header.Name = "李长全"
	header.Age = 26

	fmt.Println("第一次打印")
	printList(header)

	//第二种赋值方式
	p := NewTeacher("倪海婷", 26)
	fmt.Println("第二次打印")
	printList(p)

	p = new(Teacher)
	p.Name = "c"
	p.Age = 100
	header.Next.Next = p

	fmt.Println("第三次打印")
	printList(header)
}

func printList(h *Teacher) {
	for h != nil {
		fmt.Printf("Name:%v Age:%v\n", h.Name, h.Age)
		h = h.Next
	}
}

package main

import "fmt"

type People struct {
	Name string
	Age  int
	Next *People
}

//创建people,返回是指针
func NewPeople(name string, age int) *People {
	//new出来的就是指针类型
	p := new(People)
	p.Name = name
	p.Age = age
	return p
}

func main() {
	createList()
	//testCreateInHeader()
	//testCreateInTail()
}

func createList() {
	head := &People{}
	head.Age = 10
	head.Name = "ll"
	fmt.Println("第一次打印")
	printList(head) //ll,10

	p := NewPeople("b", 100)
	head.Next = p
	fmt.Println("第二次打印")
	printList(head) //ll,10   b,100

	p = new(People)
	p.Name = "c"
	p.Age = 100

	head.Next.Next = p

	fmt.Println("第三次打印")
	printList(head) //ll,10	b,100	c,100
}

func printList(h *People) {
	for h != nil {
		fmt.Printf("Name:%v Age:%v\n", h.Name, h.Age)
		h = h.Next
	}
}

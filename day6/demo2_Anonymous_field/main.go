package main

import "fmt"

func (p *People) Format() string {
	return fmt.Sprintf("people name=%s&age=%d", p.Name, p.Age)
}

type People struct {
	Name string
	Age  int
}

type Student struct {
	Score int
	People
}

//直接调用匿名字段内的字段,可以理解为继承(算是go语言的语法糖)
func main() {
	var s Student
	s.Age = 200
	s.Name = "abc" //相当于 s.People.Name = "abc"
	ret := s.Format()
	fmt.Println("format result:", ret)
}

package main

import "fmt"

type People struct {
	Name string
	id   int
}

type Student struct {
	People
	Score int
}

func (p *People) Format() string {
	return fmt.Sprintf("name=%s & id=%d \n", p.Name, p.id)
}

func (s *Student) Format() string {
	return fmt.Sprintf("name=%s & age=%d \n", s.Name, s.id)
}
func main() {
	//test1()
	testMethod()
}

//继承
//如果struct嵌套了匿名结构体，那么这个结构可以直接访问匿名结构体的方法，从而实现了继承。
// 如果一个struct嵌套了另一个有名结构体，那么这个模式就叫组合
//多重继承
//如果一个struct嵌套了多个结构体，从而实现了多重继承。
func test1() {
	var s Student
	//组合模式,可以直接调用,效果跟s.People.Name一样
	s.Name = "abc"
	s.People.Name = "cdg"
	s.id = 100
	fmt.Printf("%#v\n", s)
}

func testMethod() {
	var s Student
	s.id = 200
	s.People.Name = "abc"
	ret := s.People.Format()
	fmt.Println("format result:", ret)
}

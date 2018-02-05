package main

import "fmt"

type Int int

//struct是值类型
//struct在实例化的时候才会赋值
//定义的时候不会赋值
type Test struct {
	A int
}

//结构体属性可以是任意类型,自定义类型/结构体/指针 等等
type Student struct {
	Name string
	Age  int
	t    Int
	a    Test
	c    *int
}

func main() {
	//demo()
	demo2()
	//demo3()
}

func demo() {
	var s = Student{
		Name: "lcq",
		Age:  14,
		t:    10,
		a:    struct{ A int }{A: 10},
		c:    new(int),
	}
	//如果没有new一个空间,直接给指针赋值会报错
	*(s.c) = 100
	s.a.A = 10
	//{Name:lcq Age:14 t:10 a:{A:10} c:0xc04204a080}
	fmt.Printf("%+v \n", s)
	fmt.Printf("%s,%d,%d,%d,%d \n", s.Name, s.Age, s.t, s.a.A, *(s.c)) //lcq,14,10,10,100
}

//%+v 和%v有什么区别?
//%v 只会输出字段值
//%+v 会输出字段名和值 可读性更好
func demo2() {
	var s = Student{Name: "lcq", Age: 18}
	fmt.Printf("%v \n", s)  //{lcq 18 0 {0} <nil>}
	fmt.Printf("%+v \n", s) //{Name:lcq Age:18 t:0 a:{A:0} c:<nil>}

	//测试  struct是值拷贝,如果修改里面指针字段的值,会改变原值吗? -- 会
	//如果没有给c分配内存空间,会报空指针错误
	s.c = new(int)
	s2 := s
	*(s2.c) = 100
	fmt.Printf("s中c值:%d s2中c值:%d \n", *(s.c), *(s2.c)) //100 100
}

func demo3() {
	var p1 *int = new(int)
	p2 := p1
	*p2 = 100
	fmt.Printf("s1=%d\n", *p1) //100

	//new出来的是地址,所以是地址拷贝
	var p3 = new(Student)
	(*p3).Age = 100
	//p4就是指针类型,因为p3是new出来的指针
	p4 := p3
	p4.Age = 1000
	fmt.Printf("p3=%+v\n", *p3) //name=1000
}

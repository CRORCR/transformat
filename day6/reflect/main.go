package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
}

func (s *Student) Set(name string, Age int, Sex int) {
	s.Name = name
}

func (s *Student) GetName(name string) {
	s.Name = name
}

func main() {
	//testGetTypeInfo()
	testGetValuInfo()
	//testGetAllMethod()

}

func testGetTypeInfo() {
	//测试int
	var i int
	getTypeInfo(i)

	//测试struct
	var stu Student
	getTypeInfo(&stu) //kind=ptr(ptr是指针类型)

	//测试数组
	var s [5]int
	getTypeInfo(s)

	//测试切片
	var s2 []int
	getTypeInfo(s2)
}

func getTypeInfo(a interface{}) {
	typeInfo := reflect.TypeOf(a)
	kind := typeInfo.Kind()
	fmt.Println("kind of a:", kind)

	//获得方法数量
	num := typeInfo.NumMethod()
	fmt.Println("method num:", num)

	//根据名称获得方法
	method, ok := typeInfo.MethodByName("SetName")
	if !ok {
		fmt.Println("not have method SetName")
	} else {
		fmt.Println(method)
	}
	fmt.Println()
	fmt.Println()
}

func testGetValuInfo() {
	var i int = 100
	//类型转为valueof结构体
	valueInfo := reflect.ValueOf(&i)
	//通过反射改变变量的值
	valueInfo.Elem().SetInt(200)
	tmp := valueInfo.Interface()
	//这时候tmp是指针,不能使用tmp.(int)
	val := tmp.(*int)        //interface强转为int类型
	fmt.Println("val:", val) //100
	fmt.Println("val of valueInfo:", valueInfo.Elem().Int())
	fmt.Println("type:", valueInfo.Type()) //获得类型 int
	fmt.Println("kind:", valueInfo.Kind()) //int
	fmt.Println("i=", i)

	var stu Student
	valueInfo = reflect.ValueOf(stu)
	fmt.Println("type:", valueInfo.Type()) //mian.Student  类型
	fmt.Println("kind:", valueInfo.Kind()) //struct 类别

}

func getAllMethod(a interface{}) {
	typeInfo := reflect.TypeOf(a)
	num := typeInfo.NumMethod()
	for i := 0; i < num; i++ {
		//通过下标获得方法
		method := typeInfo.Method(i)
		fmt.Println(method)
	}
}

func testGetAllMethod() {
	var stu Student
	getAllMethod(&stu)
}

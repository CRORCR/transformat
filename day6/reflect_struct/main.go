package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
	Sex  int
}

func (s *Student) Set(name string, Age int, Sex int) {
	s.Name = name
	s.Age = Age
	s.Sex = Sex
}

func (s *Student) GetName(name string) {
	s.Name = name
}
func main() {
	test()
	testStruct()
}
func test() {
	//如果不是指针类型,valueof操作就可以不加elem()
	var stu Student
	valueInfo := reflect.ValueOf(stu)
	fieldNum := valueInfo.NumField()
	fmt.Println("field num:", fieldNum) //3
}
func testStruct() {
	var stu *Student = &Student{}
	stu.Set("jim", 18, 1)

	valueInfo := reflect.ValueOf(stu)
	fieldNum := valueInfo.Elem().NumField()

	fmt.Println("field num:", fieldNum)
	sexValueInfo := valueInfo.Elem().FieldByName("Sex")
	fmt.Println("sex=", sexValueInfo.Int()) //sex=1

	//修改sex值
	sexValueInfo.SetInt(100)
	fmt.Println(stu)

	//方法的调用 Call() 参数是切片([]reflect.Value)
	//这里不需要取指针,因为这个方法是指针类型student的方法
	//func (s *Student) Set() 这是student指针类型才能调用的方法
	setMethod := valueInfo.MethodByName("Set")
	fmt.Println(setMethod)

	var params []reflect.Value
	name := "Tom"
	age := 1000
	sex := 3883

	params = append(params, reflect.ValueOf(name))
	params = append(params, reflect.ValueOf(age))
	params = append(params, reflect.ValueOf(sex))

	setMethod.Call(params)
	fmt.Println(stu)
}

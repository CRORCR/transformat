package main

import "fmt"

//声明动物接口
type Animal interface {
	Eat()
	Talk()
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func main() {
	//test()
	zoo()
}

//狗调用自己的叫 吃饭
//父类animal 指向之类狗 去调用吃 叫方法
func test() {
	d := &Dog{"小黄"}
	d.Talk()
	d.Eat()
	var a Animal = d //父接口指向子类,需要地址 &
	a.Talk()
	a.Eat()
}

//动物园一群动物,一起吃 叫
func zoo() {
	//声明动物切片 里面有四个动物
	var s []Animal
	d := &Dog{"小白狗"}
	s = append(s, d)
	d = &Dog{"小花狗"}
	s = append(s, d)
	c := &Cat{"小花猫"}
	s = append(s, c)
	c = &Cat{"小白猫"}
	s = append(s, c)

	//一起吃饭,一起叫唤
	for _, v := range s {
		v.Eat()
		v.Talk()
	}

}
func (d *Dog) Eat() {
	fmt.Println(d.Name, "吃饭")
}

func (d *Dog) Talk() {
	fmt.Println(d.Name, "叫唤")
}

func (d *Cat) Eat() {
	fmt.Println(d.Name, "吃饭")
}

func (d *Cat) Talk() {
	fmt.Println(d.Name, "叫唤")
}

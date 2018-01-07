package main

import "fmt"

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

func (c *Cat) Eat() {
	fmt.Println(c.Name, " eat")
}

func (c *Cat) Talk() {
	fmt.Println(c.Name, " talk")
}

func (d *Dog) Eat() {
	fmt.Println(d.Name, " eat")
}

func (d *Dog) Talk() {
	fmt.Println(d.Name, "  talk")
}

func main() {
	//var a Animal
	//var d Dog
	//d.Eat()
	//
	//a = &d
	//a.Eat()

	test()
}
func test() {
	var list []Animal
	d := &Dog{"小黄"}
	list = append(list, d)
	d = &Dog{"小白"}
	list = append(list, d)
	c := &Cat{"小花猫"}
	list = append(list, c)

	for _, v := range list {
		v.Eat()
		v.Talk()
	}
}

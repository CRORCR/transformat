package main

import "fmt"

type People interface {
	run()
	//say()
}

type Man struct {
	Name string
}

type WoMan struct {
	Name string
}

func (m *Man) run() {
	fmt.Println(m.Name, "奔跑")
}

func (w *WoMan) run() {
	fmt.Println(w.Name, "runing")
}
func main() {
	var list []People
	m := &Man{"lcq"}
	w := &WoMan{"woman"}
	list = append(list, m)
	list = append(list, w)
	for _, value := range list {
		value.run()
	}
}

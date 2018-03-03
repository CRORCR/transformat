package main

import (
	"fmt"
	"math/rand"
	demo2_sort2 "transformat/day7/demo2_sort"
)

//定义一个学生
type Student struct {
	Name  string
	id    int
	score float32
}

//需要实现Interface三个方法,所以创建这个类型,存储学生
type studentSclice []*Student

//这里不用指针,因为指针不能找到长度,如果非要用指针,取长度就要加上*
//func (p *studentSclice) Len() int {
//	return len(*p)
//}
func (p studentSclice) Len() int {
	return len(p)
}
func (p studentSclice) Less(i, j int) bool {
	return p[i].score < p[j].score
}
func (p studentSclice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//*****
//排序肯定是一个集合才会排序,单个对象无法排序,所以排序实现的接口,需要集合去实现
func main() {
	var arr studentSclice
	for i := 0; i < 5; i++ {
		s := &Student{
			Name:  fmt.Sprintf("李%d", i),
			id:    rand.Intn(100),       //100以内
			score: rand.Float32() * 100, //100以内
		}
		arr = append(arr, s)
	}
	//排序
	//sort.Sort(arr)
	demo2_sort2.Bubble(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v", arr[i]) //这是切片,存储的是指针
	}
}

package main

import (
	"fmt"
	"math/rand"
	"oldboy/day7/demo2_sort"
)

type Student struct {
	Name  string
	id    int
	score float32
}

type studentSclice []*Student

func (p studentSclice) Len() int {
	return len(p)
}
func (p studentSclice) Less(i, j int) bool {
	return p[i].score < p[j].score
}
func (p studentSclice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var arr studentSclice
	for i := 0; i < 5; i++ {
		s := &Student{
			Name:  fmt.Sprintf("李%d", i),
			id:    rand.Intn(100),
			score: rand.Float32() * 100,
		}
		arr = append(arr, s)
	}
	//排序
	//sort.Sort(arr)
	demo2_sort.Bubble(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v", arr[i])
	}
}

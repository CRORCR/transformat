package main

import (
	"fmt"
	"sort"
)

//map初始化不会分配内存,需要make来创建内存
//map的增删改查
//map是引用类型
func main() {
	//m := make(map[int]string, 4)
	//addMap(m)
	//deleteMap(m)
	//updateMap(m)
	//selectMap(m)
	//seliceMap()
	sortMap()
}

func addMap(m map[int]string) {
	m[1] = "a"
	m[2] = "b"
	m[3] = "c"
	m[4] = "d"
	fmt.Println("add:", m)
}

func deleteMap(m map[int]string) {
	delete(m, 1)
	fmt.Println("delete:", m)
}
func updateMap(m map[int]string) {
	m[2] = "hello"
	fmt.Println("update:", m)
}
func selectMap(m map[int]string) {
	if v, ok := m[4]; ok {
		fmt.Println(v)
	}
}

func seliceMap() {
	s := make([]map[int]string, 3)
	for i := 0; i < 3; i++ {
		s[i] = make(map[int]string, 1)
		str := fmt.Sprintf("a:%d", i)
		s[i][i] = str
	}
	fmt.Println(s)
}

func rangeMap(m map[int]string) {
	for k, v := range m {
		fmt.Println("\t", k, v)
	}
}

//map排序
//先获取所有key，把key进行排序
//Map反转
//初始化另外一个map，把key、value互换即可
func sortMap(){
	m:=make(map[int]string,3)
	m[3]="33"
	m[2]="22"
	m[1]="11"
	//1.先排序
	var arrayList []int
	for v:=range m{
		arrayList=append(arrayList,v)
	}
	sort.Ints(arrayList)
	//根据健找对应值
	for _,v:=range arrayList{
		fmt.Println(v,m[v])
	}
}

//map反转
func recoverMap(){
	m:=make(map[int]string,3)
	m[3]="33"
	m[2]="22"
	m[1]="11"

	m2:=make(map[string]int,3)
	for k,v:=range m{
		m2[v]=k
	}
	fmt.Println(m2)

}
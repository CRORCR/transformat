package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{38, 1, 100, 5, 10}
	//demo1(a)
	//demo2(a)
	//demo3(a)
	demo4(a)
	fmt.Println(a) //[1 4 5 10 38]
}

//数组自有排序方法
func demo1(a []int) {
	sort.Ints(a)
}

//冒泡排序  大的往后放
func demo2(a []int) {
	for i := len(a) - 1; i > 0; i-- { //次数控制
		for j := 0; j < i; j++ { //角标值排序
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

//选择排序 小的往前放
func demo3(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
}

//插入排序---跟打扑克一样
//从第二张牌排序---从i=1开始
//从右往左,从大到小 ---j=i j--

//38
//1, 38,
//1, 4, 38,
//1, 4, 5, 38,
//1, 4, 5, 10, 38
func demo4(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break //如果发现已经放在正确的位置,就直接退出,不用比较
			}
		}
	}
}

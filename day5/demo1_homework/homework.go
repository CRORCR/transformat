package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{38, 1, 4, 5, 10}
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
			}
		}
	}
}

//快速排序  找到一个数,左边都比他小,右边都比他大
//50, 100, 4, 5, 10, 50
//10, 100, 4, 5, 38, 50
//10, 38, 4, 5, 100, 50
//10, 4, 38, 5, 100, 50
//10, 4, 5, 38, 100, 50
//5, 4, 10,
//4, 5, 10, 38,
func partion(a []int, left, right int) int {
	var i = left
	var j = right
	for i < j {
		for j > i && a[j] > a[left] {
			j--
		}
		a[j], a[left] = a[left], a[j]
		for i < j && a[i] < a[left] {
			i++
		}
		a[left], a[i] = a[i], a[left]
		fmt.Println(i)
	}
	return i
}

func qsort(a []int, left, right int) {
	//如果左边等于右边,就是只有一个元素,不需要排序,直接返回
	if left >= right {
		return
	}

	mid := partion(a, left, right)
	qsort(a, left, mid-1)
	qsort(a, mid+1, right)
}

package main

import "fmt"

func main() {
	//1.冒泡排序
	//bsort()
	//2.选择排序
	//ssort()
	//3.插入排序
	//isort()
	//4.快速排序 使用了递归
	qsort()
}

func bsort(){
	a := [...]int{8, 7, 5, 10, 15,4, 3}
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
			fmt.Println(a)
		}
	}
	fmt.Println(a)

}

//2.选择排序,选出最大或者最小的数,然后交换位置
func ssort(){
	a := [...]int{8, 7, 5, 4, 3, 10, 15}
	for i := 0; i < len(a); i++ {
		var min int = i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		//最小的元素和i位置元素互换位置
		a[i], a[min] = a[min], a[i]
	}
}

//3.插入排序
//左边是有序,右边每次过来一个元素都去有序数列进行比较,插入操作
func isort(){
	a := [...]int{8, 7, 5, 4, 3, 10, 15}
	for i:=1;i<len(a)-1;i++{
		for j:=i;j>0;j--{
			if a[j]<a[j-1]{
				a[j],a[j-1]=a[j-1],a[j]
			}
		}
	}
	fmt.Println(a)
}

//4.快速排序
//左边都比自己小,右边都比自己大
func qsort(){
	b := [...]int{8,5,4,10,15,7}
	goQsort(b[:], 0, len(b)-1)
	fmt.Println(b)
}
//8,5,4,10,15,7
func goQsort(a []int, left, right int) {
	if left >= right {
		return
	}
	val := a[left]
	k := left
	//确定val所在的位置
	for i := left + 1; i <= right; i++ {
		if a[i] < val {
			//fmt.Printf("i:%d k:%d k+1:%d \n",i,k,k+1)
			//fmt.Printf("a[k]:%d a[i]:%d a[k+1]:%d \n",a[k],a[i],a[k+1])
			a[k] = a[i]
			a[i]=a[k+1]
			k++
		}
		fmt.Println(a)
	}

	a[k] = val
	goQsort(a, left, k-1)
	goQsort(a, k+1, right)
}

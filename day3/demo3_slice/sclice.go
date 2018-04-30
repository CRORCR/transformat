package main

import (
	"fmt"
	"sort"
)

func main() {
	//testSlice()
	//testArray()
	//copyArray()
	//testModifyString()
	//sortInt()
	sortStr()
}

//1.切片指向数组第一个元素地址
func testSlice() {
	var a = [10]int{1, 2, 3, 4}

	b := a[1:5]
	fmt.Printf("%p\n", b)     //2 3 4 0
	fmt.Printf("%p\n", &a[1]) //切片指向数组第一个元素地址
}

//2.给切片添加元素如果超过长度,会新开辟空间
//注意:是与第一个元素的地址(不是数组的地址)
//append 用法:如果是切片,需要展开 s=append(array...)
func testArray() {
	var a = [5]int{1, 2, 3, 4, 5}
	s := a[1:]
	//扩容之前,切片还是指向之前的数组第一个元素
	fmt.Printf("%p %p \n", &a[1], s) //0xc042066038 0xc042066038
	s = append(s, 10)
	s = append(s, 10)
	s = append(s, 10)
	//扩容之后,新开辟的内存,与数组的地址不同
	fmt.Printf("%p %p \n", &a[1], s) //0xc042066038 0xc042070040
}

//3.切片的拷贝
//copy(s,s2)
//第二个切片拷贝到第一个切片,如果长度超了,就忽略
func copyArray() {
	var a = []int{1, 2, 3, 4, 5, 6}
	b := make([]int, 1)

	copy(b, a)
	fmt.Println(b) //1
}

//4.string
//string底层也是byte数组,也可以进行切片操作
//string是不可变的,如果需要修改,就转换成切片,然后再转成string
//rune
//这是可变的数组切片,如果是中文就是三个字节,英文就是两个字节或者一个字节
//一般如果需要处理中文,就使用rune,避免使用byte带来的转换异常
func testModifyString() {
	str := "hello 你好,世界"
	s := []rune(str)
	s[7] = '哈'
	str = string(s)
	fmt.Println(str)
}

//5.排序
//sort.Ints对整数进行排序， sort.Strings对字符串进行排序, sort.Float64s对浮点数进行排序
//sort.SearchInts(a []int, b int) 从数组a中查找b，前提是a必须有序
//sort.SearchFloats(a []float64, b float64) 从数组a中查找b，前提是a必须有序
//sort.SearchStrings(a []string, b string) 从数组a中查找b，前提是a必须有序

//对整数int排序
func sortInt() {
	//1.切片排序
	a := []int{5, 4, 3, 2, 1}
	sort.Ints(a)
	fmt.Println(a)
	//2.数组排序
	var b = [...]int{5, 4, 3, 2, 1}
	sort.Ints(b[:])
	fmt.Println(b)
}

//对string排序
func sortStr() {
	var a = [...]string{"abc", "efg", "b", "A", "eeee"}
	sort.Strings(a[:])
	fmt.Println(a)
}

//有序数组,查找元素
func testIntSearch() {
	var a = [...]int{1, 8, 38, 2, 348, 484}
	sort.Ints(a[:])
	//需要切片作为元素
	index := sort.SearchInts(a[:], 348)
	fmt.Println(index)
}

//切片：切片是数组的一个引用(长度可变)，因此切片是引用类型
//如果要切片最后一个元素去掉，可以这么写：
//Slice = slice[:len(slice)-1]
//通过make创建切片
//var slice []type = make([]type, len)
//用append内置函数操作切片(如果添加切片,需要展开)
//拷贝切片copy(s2, s1),长度越界,自动忽略

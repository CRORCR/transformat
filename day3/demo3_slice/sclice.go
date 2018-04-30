package main

import "fmt"

func main() {
	//testSlice()
	//testArray()
	//copyArray()
	testModifyString()
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

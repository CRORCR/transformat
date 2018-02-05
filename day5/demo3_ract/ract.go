package main

import "fmt"

//屏幕上一个点
type Point struct {
	x int
	y int
}

//四边形 两个点确定一个四边形
//p1和p2在内存中是连续的(同一个struct字段,在内存中是连续的)
type Ract struct {
	p1 Point
	p2 Point
}

//p1和p2内存是连续的,但是p1和p2指向的值,是不连续的
type RactA struct {
	p1 *Point
	p2 *Point
}

func main() {
	//demo1()
	demo2()
	//demo3()
}
func demo1() {
	//测试Ract中字段内存地址是否连续--连续
	var r Ract
	//p1 : x addr:0xc04200e2e0 y addr:0xc04200e2e8
	fmt.Printf("p1 : x addr:%p y addr:%p \n", &(r.p1.x), &(r.p1.y))
	//p2 : x addr:0xc04200e2f0 y addr:0xc04200e2f8
	fmt.Printf("p2 : x addr:%p y addr:%p \n", &(r.p2.x), &(r.p2.y))
	//为什么y比x大8?
	//定义的类型就是int,所以是8个字节(我计算机是64位)
}

func demo2() {
	//测试RactA两个值内存布局---连续
	//p1和p2是连续的,但是p1指向的内存,和p2指向的内存是不连续的
	var r RactA
	r.p1 = new(Point)
	r.p2 = new(Point)
	//p1:0xc04203e1b0
	//p2:0xc04203e1b8
	fmt.Printf("p1:%p p2:%p \n", &(r.p1), &(r.p2))
}
func demo3() {
	//测试RactA两个指针内存布局---不连续
	var r RactA
	r.p1 = new(Point)
	r.p2 = new(Point)
	//p1 : x addr:0xc0420100a0 y addr:0xc0420100a8
	fmt.Printf("p1 : x addr:%p y addr:%p \n", &(r.p1.x), &(r.p1.y))
	//p2 : x addr:0xc0420100b0 y addr:0xc0420100b8
	fmt.Printf("p2 : x addr:%p y addr:%p \n", &(r.p2.x), &(r.p2.y))
}

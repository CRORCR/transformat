package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	//1.素数 prime number
	//prime()
	////2. 99乘法表 multiplication table
	//table()
	////3.水仙花数 narcissus
	//narcissus()
	//4.factorial 阶乘:从1乘以2乘以3乘以4一直乘到所要求的数
	//factorial()
	//5.猜数字游戏 guess number
	guessNumber()
	//6.输出指定五星 five stars
	//fiveStars(3)
	//7.完数
	//perfect()
}

//1.素数
//除了1和本身,不能被任何数整除
func prime() {
	var n, m int
	fmt.Scanf("%d%d%s", &n, &m)
	for i := n; i < m; i++ {
		if isPrime(i) == true {
			fmt.Printf("%d\n", i)
			continue
		}
	}
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		//没有余数
		if n%i == 0 {
			return false
		}
	}
	return true
}

//2. 9*9乘法表
func table() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d	", j, i, i*j)
		}
		fmt.Println()
	}
}

//3.个位(三次方)+十位(三次方)+百位((三次方))=自己 153=1+5+3
func narcissus() {
	var n, m int
	fmt.Scanf("%d,%d", &n, &m)
	for i := n; i < m; i++ {
		//个位 十位 百位
		i := n % 10
		j := (n / 10) % 10
		k := (n / 100) % 10
		sum := i*i*i + j*j*j + k*k*k
		if n == sum {
			fmt.Println(i)
		}
	}
}

func factorial() {
	var n int
	fmt.Scanf("%d", &n)

	s := sum(n)
	fmt.Println(s)
}

//0+ 1*1
//1+ 1*2
//4+ 1*3
func sum(n int) uint64 {
	var s uint64 = 1
	var sum uint64 = 0
	for i := 1; i <= n; i++ {
		s = s * uint64(i)
		fmt.Printf("%d!=%v\n", i, s)
		sum += s
	}
	return sum
}

//5.猜数字游戏
//有个小bug,会把回车也算做输入,所以每次执行都会校验两遍,
// 所以scanf 接收的时候,就加上\n,然后忽略就好了
func guessNumber() {
	n := rand.Intn(100)
	for {
		var input int
		fmt.Scanf("%d\n", &input)
		flag := false
		switch {
		case input == n:
			fmt.Println("you are right")
			flag = true
		case input > n:
			fmt.Println("bigger")
		case input < n:
			fmt.Println("less")
		}

		if flag {
			break
		}
	}
}

//6. 写一个程序，在终端打印如下图形 five stars
//A
//AA
//AAA
//AAAA
//AAAAA
func fiveStars(n int) {
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("A")
		}
		fmt.Println()
	}
}

//7.完数
//一个数如果恰好等于它的因子之和,这个数就称为完数
//例如:6=1+2+3
//找出100以内的所有完数
func perfect() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 1; i < n+1; i++ {
		if isPerfect(i) {
			fmt.Println(i)
		}
	}
}

//能被整除就是因子
func isPerfect(n int) bool {
	var sum int = 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return n == sum
}

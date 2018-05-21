package main

import "fmt"

//continue只能跳出一层循环,如果想跳出多层循环,可以使用标签和continue配合
func main() {
	demo()
	//demo2()
}

//i 0-5
//j 0-3 5 唯独没有4
func demo() {
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}

//01234
func demo2() {
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}

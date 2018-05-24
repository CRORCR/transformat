package fifi

//1  1  2  3  5  8
func Fibo() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

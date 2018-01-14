package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"time"
)

type ReadStr struct{}

func (r *ReadStr) Read(p []byte) (n int, err error) {
	fmt.Printf("len(p)=%d  :\n", len(p))

	source := "abcdefghigklmnopq"
	for i := 0; i < 30; i++ {
		index := rand.Intn(len(source))
		p[i] = source[index]
	}
	p[30] = '\n'
	return len(p), nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var randstr = &ReadStr{}
	randReader := bufio.NewReader(randstr)
	line, prefix, _ := randReader.ReadLine()
	fmt.Printf("rand:%s prefix:%s", line, prefix)
}

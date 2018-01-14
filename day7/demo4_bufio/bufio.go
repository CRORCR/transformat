package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reads := bufio.NewReader(os.Stdin)
	line, err := reads.ReadString('\n')
	if err != nil {
		fmt.Println("read error: ", err)
	}
	fmt.Println("out", line)
}

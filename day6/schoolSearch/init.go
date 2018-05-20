package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"transformat/day6"
)

var t *trie.Trie

func LoadAllSchool() error {
	t = trie.NewTrie()
	file, err := os.Open("./data.dat")
	defer file.Close()
	var content []byte
	var data [1024]byte
	for {
		n, err := file.Read(data[:])
		if err == io.EOF {
			break
		}
		content = append(content, data[0:n]...)
	}
	lines := strings.Split(string(content), "\n")
	for _, v := range lines {
		if len(v) == 0 {
			continue
		}
		sc := strings.Split(v, " ")
		if len(sc) != 3 {
			fmt.Printf("line[%s] is not right\n", v)
			continue
		}
		var s School
		s.Name, s.City = sc[0], sc[1]
		s.Province = sc[2]
		t.Add(s.Name, &s)
	}
	return err
}

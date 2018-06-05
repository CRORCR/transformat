package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var path string = "G://read"

func main() {
	//readDoc(path)
	files, err := WalkDir(path, ".txt")
	CheckErr(err)
	for i, v := range files {
		fmt.Println(i, v)
	}
}

func CheckErr(err error) {
	if nil != err {
		panic(err)
	}
}

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}

		if fi.IsDir() { // 忽略目录
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}

		return nil
	})

	return files, err
}

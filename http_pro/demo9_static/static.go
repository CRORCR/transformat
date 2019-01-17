package main

import "net/http"

func main() {
	//http.Handle("/static/", http.FileServer(http.Dir("./day10/demo9_static/static/")))
	//会把请求中的static加到文件搜索末尾,需要删除url自带的static  真实项目,看自己需求
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./day10/demo9_static/static/"))))
	http.ListenAndServe(":8080", nil)
}

package main

import (
	"html/template"
	"net/http"
)

type People struct {
	Name string
	Age  int
}

var temp *template.Template

func init() {
	t, _ := template.ParseFiles("K:/workspace/src/transformat/day10/demo7_format/index.html")
	temp = t
}
func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	var ps []People
	for i := 0; i < 4; i++ {
		p := People{"lcq", 11}
		ps = append(ps, p)
	}
	temp.Execute(w, ps)
}

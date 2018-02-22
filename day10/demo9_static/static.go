package main

import "net/http"

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./day10/demo9_static/static/"))))
	http.ListenAndServe(":8080", nil)
}

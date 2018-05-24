package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/list/", hand)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
func hand(w http.ResponseWriter, r *http.Request) {
	//得到list后面的目录
	path := r.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		//panic(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	w.Write(all)

}

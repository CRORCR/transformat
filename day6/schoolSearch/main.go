package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func searchInQuery(querys []string) []byte {
	var result map[string]interface{} = make(map[string]interface{}, 16)
	if len(querys) == 0 {
		result["code"] = 1001
		result["message"] = "please input query"
		data, err := json.Marshal(result)
		if err != nil {
			result["code"], result["message"] = 1002, "internal 500"
		}
		return data
	}

	result["code"] = 0
	result["message"] = "success"
	var schools []*School
	searchResult := t.PrefixSearch(querys[0])
	for _, v := range searchResult {
		s, ok := v.Data.(*School)
		if !ok {
			continue
		}
		schools = append(schools, s)
	}
	result["data"] = schools
	data, err := json.Marshal(result)
	if err != nil {
		result["code"] = 1002
		result["message"] = "internal 500"
	}
	return data
}

func main() {
	//加载学校信息,放入init函数中
	err := LoadAllSchool()
	if err != nil {
		fmt.Printf("load all school failed, err:%v", err)
		return
	}

	http.HandleFunc("/search", search)
	http.ListenAndServe(":8080", nil)
}

func search(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	queryArray := r.Form["query"]

	data := searchInQuery(queryArray)
	w.Write(data)
}

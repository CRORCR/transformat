package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
//http://localhost:8080/search?query=北京交通大学
func main() {
	err:=LoadAllSchool()
	if err!=nil{
		fmt.Println("加载失败")
	}
	http.HandleFunc("/search",search)
	http.ListenAndServe(":8080",nil)
}

func search(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	queryInfo:=r.Form["query"]
	data :=searchInQuery(queryInfo)
	w.Write(data)
}

func searchInQuery(str []string)(data []byte){
	result:=make(map[string]interface{},16)
	if len(str)==0{
		result["code"]=1001
		result["meaasge"]="please input request"
		data,_=json.Marshal(result)
		return data
	}

	result["code"]=0
	result["message"]="success"

	var schools []*School
	searchResult:=t.PrefixSearch(str[0])
	for _,v :=range searchResult{
		s,ok :=v.Data.(*School)
		if !ok{
			continue
		}
		schools=append(schools,s)
	}
	result["data"]=schools
	data, _ = json.Marshal(result)
	return
}
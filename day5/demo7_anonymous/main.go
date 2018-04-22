package main

import "fmt"

type People struct{
	Score int
	Name string
}
type Student struct{
	Name string
	Age int
	People
}

func main() {
	var s Student
	s.Name="lcq"
	s.Age=26
	s.People.Score=100
	//{Name:lcq Age:26 Prople:{Score:100 Name:}}
	fmt.Printf("%+v",s)
}


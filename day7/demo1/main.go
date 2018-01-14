package main

type User struct {
	Name string
	id   int
}

type userList []*User

func (u userList) Len() int {
	return len(u)
}

func (u userList) Less(i, j int) bool {

}

func main() {

}

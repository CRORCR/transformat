package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type UserInfo struct{
	User_id int `db:"user_id"`
	User_name string `db:"user_name"`
	Sex string `db:"sex"`
	Emial string `db:"email"`
}

func main() {
	db,err:=sqlx.Open("mysql","root:root@tcp(127.0.0.1:3306)/godb")
	if err!=nil{
		fmt.Println("connection is failed")
		return
	}
	fmt.Println("connection success")
	//1.执行插入操作
	//db.Exec("insert into user_info (user_name,sex,email) values (?,?,?)","user01","男","112037@qq.com")
	//2.执行查询  get查询一条记录  select 查询集合
	var user UserInfo
	db.Get(&user,"select user_id,user_name,sex,email from user_info where user_id=?",10)
	fmt.Println(user)

	var userList []*UserInfo
	db.Select(&userList,"select user_id,user_name,sex,email from user_info where user_id>0")
	for _,v :=range userList{
		fmt.Println(v)
	}

	//3.更新表
	db.Exec("update user_info set user_name=? where user_id=?","lcq",10)

	//4.删除表
	db.Exec("delete from user_info where user_id=?",11)

	db.Close()
}
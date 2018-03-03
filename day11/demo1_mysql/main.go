package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
var Db *sqlx.DB
type UserInfo struct {
	UserId   int    `db:"user_id"`
	UserName string `db:"user_name"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

func main() {
	db, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/godb")
	if err != nil {
		fmt.Println("connection to mysql is failed,", err)
		return
	}
	fmt.Println("connection is success")

	//1.插入一条记录
	result, err := db.Exec("insert into user_info(user_name,sex,email) values(?,?,?)", "user01", "男", "112233@qq.com")
	if err != nil {
		fmt.Println("insert is failed,err:", err)
		return
	}
	userid, _ := result.LastInsertId()
	fmt.Println(userid)

	//2.查询  get查询单条记录,select 查询的结果是切片(集合)
	var userinfo UserInfo
	//java idea不支持go的第三方库,虽然显示红色,但是可以运行
	//第一个参数是结构体,第二个参数是sql语句 get只能查询一条记录
	db.Get(&userinfo, "select user_id, user_name, sex, email from user_info where user_id=2")
	fmt.Println(userinfo)

	var userList []*UserInfo
	db.Select(&userList, "select user_id, user_name, sex, email from user_info where user_id>2")
	for _, v := range userList {
		fmt.Println(v)
	}

	//更新表
	db.Exec("update user_info set user_name=? where user_id=?", "lcq01", "1")

	//删除记录
	db.Exec("delete from user_info where user_id=?", "1")
	db.Close()
}

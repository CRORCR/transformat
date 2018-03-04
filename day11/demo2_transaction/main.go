package main
import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
//为了原子性
func main() {
	db,err:=sqlx.Open("mysql","root:root@tcp(127.0.0.1:3306)/godb")
	if err!=nil{
		fmt.Println("connection is failed")
		return
	}
	//begin开启一个事物,返回一个连接和错误
	conn,_:=db.Begin()
	_,err=conn.Exec("insert into user_info (user_name,sex,email) values (?,?,?)","user01","男","112037@qq.com")
	//如果任何一个操作有问题,就要回滚
	if err!=nil{
		fmt.Println("insert is failed")
		conn.Rollback()
		return
	}
	_,err=conn.Exec("insert into user_info (user_name,sex,email) values (?,?,?)","user01","男","112037@qq.com")
	if err!=nil {
		fmt.Println("insert is failed")
		conn.Rollback()
	}
	//全部成功,提交事物
	conn.Commit()
}

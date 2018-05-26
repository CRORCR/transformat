package main

import (
	"transformat/project/logadmin/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jmoiron/sqlx"

	_ "transformat/project/logadmin/routers"
)

var Db *sqlx.DB

func main() {
	sqlLink := beego.AppConfig.String("sqlLink")
	db, err := sqlx.Open("mysql", sqlLink)
	if err != nil {
		panic(err)
	}
	Db = db
	logs.Warning("init Db success")

	models.InitDb(Db)
	beego.Run()
}

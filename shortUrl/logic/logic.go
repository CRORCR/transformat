package logic

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"transformat/shortUrl/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

//定义一个操作数据库的结构体
type ShortUrl struct {
	Id int `db:"id"`
	ShortUrl string `db:"short_url"`
	OriginUrl string `db:"origin_url"`
	HashCode string `db:"hash_code"`
}

func InitDb(dns string)(err error){
	Db,err=sqlx.Open("mysql",dns)
	if err!=nil{
		fmt.Println("连接数据库有误,err:",err)
		return
	}
	err=Db.Ping()
	return
}

func Long2Short(longRequest *model.Long2ShortUrlRequest)(response *model.Long2ShortUrlReponse,err error){
	response=&model.Long2ShortUrlReponse{}
	//1.转换MD5
	md5String:=fmt.Sprintf("%x",md5.Sum([]byte(longRequest.OriginUrl)))
	//2.查询数据库
	var shortUrl ShortUrl
	err=Db.Get(&shortUrl,"select id,short_url,origin_url,hash_code from short_url where hash_code=?",md5String)
	//2.1 如果没有
	if err==sql.ErrNoRows{
		//ErrNoRows 没有记录,不算是错误
		err=nil
		short,er:=getShortUrl(longRequest,md5String)
		if er==nil{
			err=er
			return
		}
		response.ShortUrl=short
		return
	}
	//2.2如果不是找不到记录的错误
	if err!=nil{
		return
	}
	//3.如果数据库存在,就返回
	response.ShortUrl=shortUrl.ShortUrl
	return
}

func getShortUrl(longRequest *model.Long2ShortUrlRequest,MD5string string)(shortUrl string,err error){
	//开启一个事物
	conn,_:=Db.Begin()
	result,err:=conn.Exec("insert into short_url (origin_url,hash_code) values(?,?)",longRequest.OriginUrl,MD5string)
	if err!=nil{
		conn.Rollback()
		return
	}
	id,_:=result.LastInsertId()
	shortUrl=transTo62(id)
	_,err=conn.Exec("update short_url set short_url=? where id=?",shortUrl,id)
	conn.Commit()
	return
}

func transTo62(id int64)(short string){
	charSet:="0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var str []byte
	for{
		str=append(str,charSet[id%62])
		id=id/62
		if id==0{
			break
		}
	}
	short=string(str)
	return
}

func Short2Long(shortReq *model.Short2LongRequest)(response *model.Short2LongUrlReponse,err error){
//1.根据短url找长url
	response=&model.Short2LongUrlReponse{}
	var short ShortUrl
	//数据库根据短地址查找
	err=Db.Get(&short,"select id,short_url,origin_url,hash_code from short_url where short_url=?",shortReq.ShortUrl)
	//查不到,跳转404
	if err==sql.ErrNoRows{
		response.Code=404
		return
	}
	//2.找不到就报错
	if err!=nil {
		response.Code = 500 //内部错误
		return
	}
	//3.找到就返回
	response.OriginUrl=short.OriginUrl
	return
}

func GetLastUrlTop10(limit int)(result []*model.ShortUrl,err error){
	err=Db.Select(&result,"select short_url from short_url order by id desc limit?",limit)
	return
}
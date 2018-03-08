package logic

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"transformat/short_url/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var(
	Db *sqlx.DB
)

func InitDb(dns string)(err error){
	Db,err=sqlx.Open("mysql",dns)
	if err!=nil{
		fmt.Println("connection is failed,",err)
		return
	}
	return
}
//定义一个操作数据库的结构体
type ShortUrl struct {
	Id int `db:"id"`
	ShortUrl string `db:"short_url"`
	OriginUrl string `db:"origin_url"`
	HashCode string `db:"hash_code"`
}

//先查一下,对应的短地址在数据库是否存在,如果存在就返回
//如果不存在,就用自增id作为短地址,这个id转为64进制返回
func Long2Short(req *model.Long2ShortRequest)(reponse *model.Long2ShortResponse,err error){
	reponse=&model.Long2ShortResponse{}

	//长地址使用md5加密后,是16长度的数组
	// 转换成字符串,去数据库查询
	urlMd5:=fmt.Sprintf("%x",md5.Sum([]byte(req.OriginUrl)))

	var short ShortUrl
	err=Db.Get(&short,"select id,short_url,origin_url,hash_code from short_url where hash_code=?",urlMd5)
	//sql.ErrNoRows 这是sql自己的错误码,表示没有记录
	fmt.Println("111",err)
	if err==sql.ErrNoRows{
		err=nil
		//如果没有记录,就自己生成一个短url
		shortUrl,er:=generateShortUrl(req,urlMd5)
		if er!=nil{
			err=er
			return
		}
		reponse.ShortUrl=shortUrl
		fmt.Println("reponse.ShortUrl:",reponse.ShortUrl)
		return
	}

	//如果还是有错误,就直接返回
	if err!=nil{
		return
	}
	//数据库的短url返回
	reponse.ShortUrl=short.ShortUrl
	return
}

//生成短url
func generateShortUrl(req *model.Long2ShortRequest,hashCode string)(shortUrl string,err error){

	conn,_:=Db.Begin()
	//数据库插入一条记录
	result,err:=conn.Exec("insert into short_url (origin_url,hash_code) values(?,?)",req.OriginUrl,hashCode)
	if err!=nil{
		conn.Rollback()
		return
	}
	//得到自增的id,转成62进制,这样 0-9a-zA-Z 这62个字符都利用上了,节约空间
	insertId,_:=result.LastInsertId()
	shortUrl=transTo62(insertId)

	//存储到数据库,并返回
	_,err=conn.Exec("update short_url set short_url=? where id=?",shortUrl,insertId)
	conn.Commit()
	return
}

//十进制转62进制   %62结果
//1%62=1  10%62=a 61%62=Z 62%62=10
func transTo62(id int64)(str string){
	charSet:="0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var shortUrl []byte
	for{
		number := id%62
		shortUrl=append(shortUrl,charSet[number])
		id=id/62
		if id==0{
			break
		}
	}
	return string(shortUrl)
}
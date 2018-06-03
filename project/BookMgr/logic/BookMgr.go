package logic

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/garyburd/redigo/redis"
	"io/ioutil"
	"sync"
	"time"
)

type BookMgr struct {
	//所有的图书集合
	BookList []*Book
	//map 所有作者的书
	AuthorList map[string][]*Book
	//map 同名所有的书
	NameList map[string][]*Book
	//map 图书被谁所借
	BookForStu map[int][]*Student
	//锁
	lock sync.Mutex
}

func NewMgr() (bookMgr *BookMgr) {
	bookMgr = &BookMgr{
		AuthorList: make(map[string][]*Book, 16),
		NameList:   make(map[string][]*Book, 16),
		BookForStu: make(map[int][]*Student, 16),
	}
	return
}

//添加一本书 从数据库查询,如果没有就更新一条,如果有的话就增加num
func (bm *BookMgr) AddBook(book *Book) {
	bm.BookList = append(bm.BookList, book)
	var bookTmp Book
	err := Db.Get(&bookTmp, "select bookId,num,name,author,publishDate from book where name=? and author=?", book.Name, book.Author)
	//如果有错误,直接返回
	if err != nil && err != sql.ErrNoRows {
		return
	}
	//如果没有查到记录,就插入一条
	if err == sql.ErrNoRows {
		_, err = Db.Exec("insert into book(num,name,author,publishDate)values(?,?,?,?)", book.Num, book.Name, book.Author, book.PublishDate)
	} else {
		_, err = Db.Exec("update book set num = num + ? where name=? and author=?",
			book.Num, book.Name, book.Author)
	}
	return
}

//保存一本书到本地磁盘
func (bm *BookMgr) save() {
	data, err := json.Marshal(bm)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("d:/book.json", data, 0666)
	if err != nil {
		fmt.Printf("write file failed, err:%v\n", err)
		return
	}
}

//借书
func (bm *BookMgr) Borrow(stu *Student, bookId int) (err error) {
	bm.lock.Lock()
	defer bm.lock.Unlock()
	//是否存在这本书
	var bookTmp *Book
	for _, v := range bm.BookList {
		if v.BookId == bookId {
			bookTmp = v
			break
		}
	}
	if bookTmp == nil {
		err = fmt.Errorf("book id[%s] it not exist", bookId)
		return
	}
	err = bookTmp.Borrow(1)
	if err == nil {
		return
	}
	stu.AddBook(bookTmp)
	//保存到本地
	bm.save()
	return
}

func (b *BookMgr) SearchByBookName(bookName string) (bookList []*Book) {
	sql := fmt.Sprintf("select bookId, name, author, num, publishDate from book where name like '%%%s%%'", bookName)
	Db.Select(&bookList, sql)
	return
}
func (b *BookMgr) SearchByAuthor(Author string) (bookList []*Book) {
	sql := fmt.Sprintf("select bookId, name, author, num, publishDate from book where author like '%%%s%%'", Author)
	fmt.Println(sql)

	Db.Select(&bookList, sql)
	return
}

func (b *BookMgr) SearchByPushlish(min time.Time, max time.Time) (bookList []*Book) {
	sql := fmt.Sprintf("select bookId, name, author, num, publishDate from book where publishDate > ? and publishDate < ?",
		min, max)
	Db.Select(&bookList, sql)
	return
}

func (b *BookMgr) GetTop10() (bookList []*Book) {
	sql := fmt.Sprintf("select bookId, name, author, num, publishDate, borrowCount from book order by borrowCount desc limit 10")
	Db.Select(&bookList, sql)
	return
}

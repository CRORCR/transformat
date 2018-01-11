package login

import (
	"errors"
	"sync"
)

//书籍信息包括书名、副本数、作者、出版日期
type Book struct {
	Bookid int
	Name   string
	Num    int
	Author string
	data   int64
	lock   sync.Mutex
}

//书籍录入功能  返回一本书
func NewBook(id int, name string, num int, author string, data int64) (book *Book) {
	book = &Book{Bookid: id, Name: name, Num: num, Author: author, data: data}
	return
}

//借书
func (b *Book) Borrow() (err error) {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.Num <= 0 {
		err = errors.New("图书已经被借完~~~")
		return
	}
	b.Num -= 1
	return
}

//还书
func (b *Book) Back() (err error) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.Num += 1
	return
}

//书籍查询功能，按照书名、作者、出版日期等条件检索

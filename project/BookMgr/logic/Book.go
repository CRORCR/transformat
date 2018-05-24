package logic

import (
	"fmt"
	"sync"
	"time"
)

type Book struct {
	//图书编号
	BookId int `json:"bookId"`
	Num    int `json:"num"`
	//被借过数量
	Count int `json:"count"`
	//图书名称
	Name string `json:"name"`
	//作者
	Author string `json:"author"`
	//出版日期
	PublishDate time.Time `json:"publishData"`
	//锁
	lock sync.Mutex
}

//创建一本书
func NewBook(bookId, num int, name, author string, data time.Time) (book *Book) {
	book = &Book{BookId: bookId, Num: num, Name: name, Author: author, PublishDate: data}
	return
}

//提供借书函数
func (b *Book) Borrow(num int) (err error) {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.Num < num {
		return fmt.Errorf("error:%v \n", err)
	}
	//图书数量 -1
	//被借过的图书总数 +1
	b.Count = b.Count + num
	b.Num = b.Num - num
	return
}

func (b *Book) Back(num int) {
	b.lock.Lock()
	defer b.lock.Unlock()

	b.Num = b.Num + num
}

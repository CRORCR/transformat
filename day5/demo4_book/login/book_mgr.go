package login

import (
	"errors"
	"fmt"
	"sync"
)

type Bookmgr struct {
	//图书集合
	BookList []*Book
	//图书名-->学生集合
	BookStudentMap map[int][]*Student
	//书名-->书集合
	BookNameMap map[string][]*Book
	//作者-->书集合
	BookAuthorMap map[string][]*Book
	lock          sync.Mutex
}

//创建图书管理
func NewBookMgr() (b *Bookmgr) {
	b = &Bookmgr{
		BookStudentMap: make(map[int][]*Student, 16),
		BookNameMap:    make(map[string][]*Book, 16),
		BookAuthorMap:  make(map[string][]*Book, 16),
	}
	return
}

//增加一本书
func (bm *Bookmgr) Add(b *Book) (err error) {
	bm.lock.Lock()
	defer bm.lock.Unlock()
	//1.添加到管理系统 集合中
	bm.BookList = append(bm.BookList, b)

	//2.更新书名-->书集合
	bookList, ok := bm.BookNameMap[b.Name]
	if !ok {
		var tem []*Book
		tem = append(tem, b)
		bm.BookNameMap[b.Name] = tem

	} else {
		bookList = append(bookList, b)
		bm.BookNameMap[b.Name] = bookList
	}

	//更新 作者-->书集合
	authorList, ok := bm.BookAuthorMap[b.Author]
	if !ok {
		var tem []*Book
		tem = append(tem, b)
		bm.BookAuthorMap[b.Author] = tem
	} else {
		authorList = append(authorList, b)
		bm.BookAuthorMap[b.Author] = authorList
	}
	return
}

//	BookNameMap map[string][]*Book
//根据书名查询
func (bm *Bookmgr) FindbyBookName(name string) (list []*Book) {
	bm.lock.Lock()
	defer bm.lock.Unlock()
	list = bm.BookNameMap[name]
	return
}

//根据作者查询
//BookAuthorMap map[string][]*Book
func (bm *Bookmgr) FindByAuthor(author string) (list []*Book) {
	bm.lock.Lock()
	defer bm.lock.Unlock()
	list = bm.BookAuthorMap[author]
	return
}

//根据出版日期  这是一个区间值
func (bm *Bookmgr) FindBydata(min, max int64) (list []*Book) {
	bm.lock.Lock()
	defer bm.lock.Unlock()
	for _, v := range bm.BookList {
		if v.data < min || v.data > max {
			continue
		}
		list = append(list, v)
	}
	return
}

//借书
func (bm *Bookmgr) Borrow(s *Student, b *Book) (err error) {
	bm.lock.Lock()
	bm.lock.Unlock()
	var book *Book
	for _, v := range bm.BookList {
		//有没有这本书
		if v.Bookid == b.Bookid {
			book = v
			break
		}
	}
	if book == nil {
		err = errors.New("不存在此书")
	}

	s.AddBook(book)

	err = b.Borrow()
	if err != nil {
		return
	}
	return
}

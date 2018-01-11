package login

import (
	"fmt"
	"sync"
)

//学生信息管理:姓名、年级、身份证、性别、 借了什么书等信息
type Student struct {
	Id       int
	Name     string
	Grade    int
	Identify string
	Sex      int
	BookMap  map[int]*Book
	lock     sync.Mutex
}

//创建一个学生
func NewStudent(id int, name string, grade int, identify string, sex int) (s *Student) {
	s = &Student{Id: id, Name: name, Grade: grade, Identify: identify, Sex: sex, BookMap: make(map[int]*Book, 16)}
	return
}

//增加一本书
func (s *Student) AddBook(b *Book) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.BookMap[b.Bookid] = b
	return
}

//减少一本书
func (s *Student) BackBook(bid int) (err error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	//判断是否存在map中,存在,删除,不存在 报错
	_, ok := s.BookMap[bid]
	if !ok {
		err = fmt.Errorf("%s学生没有借过%d", s.Name, bid)
	}
	delete(s.BookMap, bid)
	return
}

//借过多少书 集合
func (s *Student) GetList() (list []*Book) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, v := range s.BookMap {
		list = append(list, v)
	}
	return
}

图书馆借书
会存在并发,需要加锁,使用defer


学生实体里面,map[string]Book   key是string,为什么赋值bookid 这个值是int型
book类,还书为什么需要加锁
查找书籍为什么要加锁


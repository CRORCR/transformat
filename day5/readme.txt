Go中的struct
1.用来自定义复杂数据结构
2.struct⾥里面可以包含多个字段（属性），字段可以是任意类型(自定义类型/结构体/指针 等等)
3.struct类型是值类型
4.struct类型可以嵌套
5.Go语言没有class类型，只有struct类型
6.struct在实例化的时候才会赋值,定义的时候不会赋值

%+v 和%v有什么区别?
%v 只会输出字段值
%+v 会输出字段名和值 可读性更好

struct定义的三种形式：
1. var stu Student
2. var stu *Student = new (Student)
3. var stu *Student = &Student{}
//第一种的创建一个值类型的struct
//第二和第三都是返回一个指针

其中2和3返回的都是指向结构体的指针，访问形式如下：
stu.Name 或者 (*stu).Name 效果一样,前者go进行了自动类型转换,可以不用先取地址,再去调用

//函数和方法?
go语言中,定义在struct结构体中的就是方法
全局的就是函数

struct的初始化
struct的内存布局：struct中的所有字段在内存是连续的
每个节点包含下一个节点的地址，这样把所有的节点串串起来了，通常把
链表中的第一个节点叫做链表头

双链表定义
type Student struct {
    Name string
    Next* Student
    Prev* Student
}
如果有两个指针分别指向前一个节点和后一个节点，我们叫做双链表

二叉树定义
type Student struct {
    Name string
    left* Student
    right* Student
}

如果每个节点有两个指针分别⽤用来指向左子树和右子树，我们把这样的
结构叫做二叉树

工厂模式
golang中没有构造方法,一般可以使用工厂方式来初始化struct
Package model
type student struct {
    Name stirng
    Age int
}
func NewStudent(name string, age int) *student {
    return &student{
                Name:name,
                Age:age,
           }
}
Package main
S := new (student)
S := model.NewStudent(“tony”, 20)


1. make用来分配map、slice、channel类型的内存
2. new用来分配值类型的内存

struct中的tag
我们可以为struct中的每个字段，写上一个tag。这个tag可以通过反射的
机制获取到，最常用的场景就是json序列化和反序列化

type student struct {
    Name stirng `json=“name”`
    Age int `json=“age”`
}

匿名字段
结构体中字段可以没有名字，即匿名字段
type Car struct {
    Name stirng
    Age int
}
type Train struct {
    Car
    Start time.Time
    int
}

继承
如果⼀一个struct嵌套了另一个匿匿名结构体，那么这个结构可以直接访问
匿名结构体的方法，从而实现了了继承。
如果⼀一个struct嵌套了另一个有名结构体，那么这个模式就叫组合

多重继承
如果⼀一个struct嵌套了了多个匿名结构体，那么这个结构可以直接访问
多个匿名结构体的方法，从而实现了多重继承。


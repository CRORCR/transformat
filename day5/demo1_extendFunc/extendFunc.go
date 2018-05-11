package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//继承方法
//struct里面只要包含了父类,就可以使用父类的方法和参数
//继承和组合其实一个一个,就是有名匿名的区别
//匿名--可以直接使用--继承
//有名--加上名称使用--组合

//多重继承--嵌套了多个匿名结构,使用它们方法和参数

//实现tostring方法,在go里面只需要提供一个string方法就可以
type Car struct {
	weight int
	name   string
}

func (p *Car) Run() {
	fmt.Println("running")
}

//实现tostring方法名必须是:String
func (p *Bike) String()string{
	str := fmt.Sprintf("name=[%s] weight=[%d]", p.name, p.weight)
	return str
}

//匿名继承,直接使用父类的方法和字段
type Bike struct {
	Car
	lunzi int
}

//有名继承,调用就需要加上字段名(c)
type Train struct {
	c Car
}

func main() {
	read:=bufio.NewReader(os.Stdin)
	str,_:=read.ReadString('\n')
	str=strings.Replace(str,"\n","",-1)
	result:=strings.Split(str," ")
	for i:=len(result)-1;i>=0;i--{
		fmt.Printf("%v ",string(result[i]))
	}

















}

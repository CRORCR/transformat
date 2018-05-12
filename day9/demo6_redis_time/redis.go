package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)
//过期时间
func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	defer c.Close()
	//先设置一个值,再设置过期时间 这里使用expire关键字, 健abc 过期时间10秒
	//10秒后就获取不到健abc对应的值
	_, err = c.Do("expire", "abc", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
}
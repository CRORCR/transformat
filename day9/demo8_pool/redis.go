package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)
//声明一些全局变量
var (
	pool          *redis.Pool
	redisServer   = flag.String("redisServer", ":6379", "")
	redisPassword = flag.String("redisPassword", "123456", "")
)
//初始化一个pool
func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		//最大空闲数
		MaxIdle:     64,
		//活跃连接数 0:表示没有限制
		MaxActive:   1000,
		//超时时间
		IdleTimeout: 240 * time.Second,
		//达到最大连接是否阻塞 false:表示返回错误
		Wait:true,
		Dial: func() (redis.Conn, error) {
			//Dial 创建链接
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			/*
            if _, err := c.Do("AUTH", password); err != nil {
                c.Close()
                return nil, err
            }*/
			return c, err
		},
		//TestOnBorrow 测试连接是否可用
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

type Student struct {
	Id int 	`json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Score float32 `json:"score"`
}

func main() {
	flag.Parse()
	pool = newPool(*redisServer, *redisPassword)

	conn := pool.Get()
	defer conn.Close()

	var stu Student = Student{
		Id: 1000,
		Name:"abc",
		Age:89,
		Score:99.2,
	}

	data, _ := json.Marshal(stu)
	//redis操作
	v, err := conn.Do("SET", 1000, string(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
	ret, err := redis.String(conn.Do("GET", 1000))
	if err != nil {
		fmt.Println(err)
		return
	}

	var stu01 Student
	json.Unmarshal([]byte(ret), &stu01)
	fmt.Printf("stu01:%#v\n", stu01)

}
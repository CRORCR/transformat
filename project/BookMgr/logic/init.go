package logic

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

var (
	Db   *sqlx.DB
	pool *redis.Pool
)

func newPool(addr, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
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
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func InitRedis(addr string, pass string) (err error) {
	pool = newPool(addr, pass)
	return
}

func InitDb(dns string) (err error) {
	Db, err = sqlx.Open("mysql", dns)
	if err != nil {
		fmt.Println("connect to msyql failed, ", err)
		return
	}
	return
}

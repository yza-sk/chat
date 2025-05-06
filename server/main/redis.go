package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

// 定义一个全局poll
var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) *redis.Pool {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", 123456); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
	return pool
}

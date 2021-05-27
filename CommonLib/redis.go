package commonlib

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var (
	ip = "10.0.72.202:6379"
)

func Set(key string, value string) {
	c, err := redis.Dial("tcp", ip)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", key, value)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	return
}

func get(key string, value string) {
	c, err := redis.Dial("tcp", ip)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("GET", key, value)
	if err != nil {
		fmt.Println("redis get failed:", err)
	}
	return
}

package redis

import (
	"github.com/go-redis/redis/v7"
	"OrderMeal/data/log"
	"OrderMeal/conf"
	"strconv"
	"time"
)

func C() *redis.Client {
	d, _ := strconv.Atoi(conf.Cfg("redis", "db"))
	client := redis.NewClient(&redis.Options{
		Addr: 			conf.Cfg("redis", "addr"),
		Password: 	conf.Cfg("redis", "pass"),
		DB:					d,
	})
	return client
}

// 设置，键-值-有效期
func Set(key, val string, sec time.Duration) bool {
	err := C().Set(key, val, sec*time.Second).Err()
	if err != nil {
		log.Error("set redis key err: ", err)
		return false
	}
	return true
}

// 读取，键
func Get(key string) (string, error) {
	val, err := C().Get(key).Result()
	if err != nil {
		log.Error("get redis value err: ", err)
		return "", err
	}
	return val, nil
}

// val:= Set("kkk", "aaaaa", 100)
// val, err := Get("kkk")
// fmt.Println("val: ", val)
// fmt.Println("err: ", err)
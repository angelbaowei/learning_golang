package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// 初始化连接  新版本V8
//func initClient() (err error) {
//	rdb = redis.NewClient(&redis.Options{
//		Addr:     "localhost:16379",
//		Password: "",  // no password set
//		DB:       0,   // use default DB
//		PoolSize: 100, // 连接池大小
//	})
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	_, err = rdb.Ping(ctx).Result()
//	return err
//}


func main() {
	res := initClient()
	fmt.Println(res)
}
package internal

//
//import (
//	"strconv"
//	"time"
//
//	"github.com/go-redis/redis"
//	// "sync"
//)
//
//var (
//	client *redis.Client
//)
//
//// RedisConfig ....
//type RedisConfig struct {
//	Addr     string `json:"Addr"`
//	DB       int    `json:"DB"`
//	Pwd      string `json:"Pwd"`
//	PoolSize int    `json:"PoolSize"`
//}
//
//// ConnectRedis ...
//func ConnectRedis(rc *RedisConfig) {
//	client = redis.NewClient(&redis.Options{
//		Addr:     rc.Addr,
//		Password: rc.Pwd,
//		DB:       rc.DB,
//		PoolSize: rc.PoolSize,
//	})
//}
//
//// EXPIRE 1 hour
//const EXPIRE = 3600 * 1
//
//// setURLCache set url and id into cache
//func setURLCache(url string, id int64) error {
//	// mutex.Lock()
//	// defer mutex.Unlock()
//	if err := client.Set(url, id, EXPIRE*time.Second).Err(); err != nil {
//		return err
//	}
//	return nil
//}
//
//// checkURLCacheExist ...
//func checkURLCacheExist(url string) (int64, bool) {
//	var val string
//	var err error
//
//	// mutex.Lock()
//	// defer mutex.Unlock()
//
//	if val, err = client.Get(url).Result(); err != nil || val == "" {
//		return -1, false
//	}
//	id, _ := strconv.ParseInt(val, 10, 64)
//	return id, true
//}
//
//// delURLCache ...
//func delURLCache(url string) error {
//	// mutex.Lock()
//	// defer mutex.Unlock()
//	_, err := client.Del(url).Result()
//	return err
//}

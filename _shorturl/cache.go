package shorturl

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

var client *redis.Client

func ConnectRedis(rc *RedisConfig) {
	client = redis.NewClient(&redis.Options{
		Addr:     rc.Addr,
		Password: rc.Pwd,
		DB:       rc.DB,
		PoolSize: rc.PoolSize,
	})
}

const EXPIRE = 3600 * 1

func SetUrlCache(url string, id int64) error {
	if err := client.Set(url, id, EXPIRE*time.Second).Err(); err != nil {
		return err
	}
	return nil
}

func CheckUrlCacheExist(url string) (int64, bool) {
	var val string
	var err error

	if val, err = client.Get(url).Result(); err != nil || val == "" {
		return -1, false
	}
	id, _ := strconv.ParseInt(val, 10, 64)
	return id, true
}

func DelUrlCache(url string) error {
	_, err := client.Del(url).Result()
	return err
}

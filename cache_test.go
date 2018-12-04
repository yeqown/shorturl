package shorturl

import (
	"testing"
)

func Test_SetURLCache(t *testing.T) {
	redisC := &RedisConfig{
		Addr:     "localhost:6379",
		Pwd:      "",
		DB:       2,
		PoolSize: 5,
	}

	ConnectRedis(redisC)
	url := "http://www.baidu.com"
	if err := setURLCache(url, 2); err != nil {
		t.Error(err)
		t.FailNow()
	}

	if id, b := checkURLCacheExist(url); !b || id == -1 {
		t.Error("Did not find url cache")
		t.FailNow()
	}

	// exist cache
	if err := delURLCache(url); err != nil {
		t.Error("Del cache err: ", err)
		t.Fail()
	}

}

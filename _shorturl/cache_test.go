package shorturl

import (
	"testing"
)

func Test_SetUrlCache(t *testing.T) {
	LoadConfig("./config.json")
	ins := GetInstance()
	ConnectRedis(ins.Redis)
	url := "http://www.baidu.com"
	if err := SetUrlCache(url, 2); err != nil {
		t.Error(err)
		t.FailNow()
	}

	if id, b := CheckUrlCacheExist(url); !b || id == -1 {
		t.Error("Did not find url cache")
		t.FailNow()
	}

	// exist cache
	if err := DelUrlCache(url); err != nil {
		t.Error("Del cache err: ", err)
		t.Fail()
	}

}

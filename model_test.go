package shorturl

import (
	"testing"
)

func Test_Insert(t *testing.T) {
	// LoadConfig("./config.json")
	ConnectDB("yeqown:yeqown@/shorturl")
	db, _ := GetDB()
	um := &URLModel{
		DB:      db,
		LongURL: "http://www.baidu.com",
	}
	if id, err := um.insert(); err != nil || id == -1 {
		t.Errorf("get err: %s, id: %d\n", err.Error(), id)
		t.FailNow()
	}
}

func Test_QueryUrl(t *testing.T) {
	ConnectDB("yeqown:yeqown@/shorturl")
	db, _ := GetDB()

	um := &URLModel{
		DB:      db,
		LongURL: "http://google.com",
	}
	um.insert()
	um.clear()

	if err := um.query(); um == nil || err != nil {
		t.Errorf("got an err: %s, and LongUrlRecord got: %T", err, um)
		t.FailNow()
	}
	if um.LongURL != "http://google.com" {
		t.Error("Query result not right")
		t.FailNow()
	}
}

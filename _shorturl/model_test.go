package shorturl

import (
	"testing"
)

func Test_Insert(t *testing.T) {
	LoadConfig("./config.json")
	ConnectDB(GetInstance().MySql)
	db, _ := GetDB()
	um := &UrlModel{
		DB:      db,
		LongUrl: "http://www.baidu.com",
	}
	if id, err := um.Insert(); err != nil || id == -1 {
		t.Errorf("get err: %s, id: %d\n", err.Error(), id)
		t.FailNow()
	}
}

func Test_QueryUrl(t *testing.T) {
	LoadConfig("./config.json")
	ConnectDB(GetInstance().MySql)
	db, _ := GetDB()

	um := &UrlModel{
		DB:      db,
		LongUrl: "http://google.com",
	}
	um.Insert()
	um.Clear()

	if err := um.Query(); um == nil || err != nil {
		t.Error("got an err: %s, and LongUrlRecord got: %t", err, um)
		t.FailNow()
	}
	if um.LongUrl != "http://google.com" {
		t.Error("Query result not right")
		t.FailNow()
	}
}

func Test_Update(t *testing.T) {

}

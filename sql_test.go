package shorturl

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_ConnectDB(t *testing.T) {
	// _ = LoadConfig("./config.json")
	// ins := GetInstance()
	if err := ConnectDB("yeqown:yeqown@/shorturl"); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func Test_GetDB(t *testing.T) {
	// _ = LoadConfig("./config.json")
	// ins := GetInstance()
	_ = ConnectDB("yeqown:yeqown@/shorturl")
	if db, err := GetDB(); db == nil || err != nil {
		t.Error("got err: ", err)
		t.FailNow()
	}
}

func Test_CloseConnection(t *testing.T) {
	_ = ConnectDB("yeqown:yeqown@/shorturl")

	CloseConnection()
}

package shorturl

import (
	"testing"
)

func Test_ConnectDB(t *testing.T) {
	_ = LoadConfig("./config.json")
	ins := GetInstance()
	if err := ConnectDB(ins.MySql); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func Test_GetDB(t *testing.T) {
	_ = LoadConfig("./config.json")
	ins := GetInstance()
	_ = ConnectDB(ins.MySql)
	if db, err := GetDB(); db == nil || err != nil {
		t.Error("got err: ", err)
		t.FailNow()
	}
}

func Test_CloseConnection(t *testing.T) {
	_ = LoadConfig("./config.json")
	ins := GetInstance()
	_ = ConnectDB(ins.MySql)

	CloseConnection()
}

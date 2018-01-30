package shorturl

import (
	"testing"
)

func Test_LoadConfig(t *testing.T) {
	if err := LoadConfig("./config.json"); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func Test_GetInstance(t *testing.T) {
	_ = LoadConfig("./config.json")
	if ins := GetInstance(); ins == nil {
		t.Error("got null config pointer")
		t.FailNow()
	}
}

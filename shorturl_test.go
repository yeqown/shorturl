package shorturl

import (
	"strings"
	"testing"
)

func Test_ShortenAndParse(t *testing.T) {
	url := "http://www.baidu.com"
	short, err := Shorten(url)
	if err != nil || short == "" {
		t.Errorf("Shorten got err: %v", err)
		t.FailNow()
	}

	arr := strings.Split(short, "/")
	long, err := Parse(arr[len(arr)-1])
	if err != nil {
		t.Error("Parse got err:", err)
		t.FailNow()
	}

	if long != url {
		t.Errorf("Parse `%s` result `%s` didn't not equal 2 origin `%s`", short, long, url)
		t.Fail()
	}
}

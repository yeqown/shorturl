package shorturl

import (
	"strings"
	"testing"
)

func Test_ShortenAndParse(t *testing.T) {
	LoadConfig("./config.json")

	url := "http://www.baidu.com"
	shorturl, err := Shorten(url)
	if err != nil || shorturl == "" {
		t.Errorf("Shorten got err:", err)
		t.FailNow()
	}

	splited := strings.Split(shorturl, "/")

	longurl, err := Parse(splited[len(splited)-1])
	if err != nil {
		t.Error("Parse got err:", err)
		t.FailNow()
	}

	if longurl != url {
		t.Errorf("Parse `%s` result `%s` didn't not equal 2 origin `%s`", shorturl, longurl, url)
		t.Fail()
	}
}

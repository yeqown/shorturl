package shorturl

import (
	"testing"
)

func Test_Encode(t *testing.T) {
	if ch := Encode(10); ch != "a" {
		t.Errorf("encode `%d` got `%s`", 10, ch)
		t.FailNow()
	}
	if ch := Encode(63); ch != "11" {
		t.Errorf("encode `%d` got `%s`", 63, ch)
		t.FailNow()
	}
}

func Test_Decode(t *testing.T) {
	if i := Decode("S"); i != 54 {
		t.FailNow()
	}
	if i := Decode("0"); i != 0 {
		t.Errorf("decode `%s` got `%d`", "0", i)
		t.FailNow()
	}
	if i := Decode("0Z"); i != 61 {
		t.Errorf("decode `%s` got `%d` actual `%d`", "0Z", i, 61)
		t.FailNow()
	}
	if i := Decode("Z0"); i != 3782 {
		t.Errorf("decode `%s` got `%d` actual `%d`", "0Z", i, 3782)
		t.FailNow()
	}
}

func Test_ProcessEncodeThenDecode(t *testing.T) {
	var i int64 = 12312
	s := Encode(i) // 3cA -- base62
	if s != "3cA" {
		t.Errorf("encode `%d` got `%s`, actual `%s`", i, s, "3cA")
		t.FailNow()
	}
	if di := Decode(s); di != i {
		t.Errorf("decode `%s` got `%d`, actual `%d`", s, di, i)
		t.FailNow()
	}
}

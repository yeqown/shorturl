package shorturl

import (
	"testing"
)

func Test_RegRouter(t *testing.T) {
	r := RegRouter()
	if r == nil {
		t.FailNow()
	}
}

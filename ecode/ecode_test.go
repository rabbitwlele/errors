package ecode

import (
	"testing"
)

func Test_New(t *testing.T) {
	code := New(100, "test")
	if code.Code() != 100 {
		t.Logf("code should be 100")
		t.FailNow()
	}
	if code.Message() != "test" {
		t.Logf("message should be 101")
		t.FailNow()
	}
}

func Test_New_Panic(t *testing.T) {
	defer func() {
		str := recover()
		if str != "code: 101 already exist" {
			t.Logf("repeating new code should be panic")
		}
	}()
	New(101, "test")
	New(101, "test")
}

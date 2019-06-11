package _0_valid_parentheses

import "testing"

func Test1(t *testing.T) {
	if isValid("[]") == false {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	if isValid("[]()[)") == true {
		t.Fail()
	}
}

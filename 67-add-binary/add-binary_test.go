package _7_add_binary

import "testing"

func Test1(t *testing.T) {
	a := "11"
	b := "1"
	expect := "100"
	if expect != addBinary(a, b) {
		t.Fail()
	}
}

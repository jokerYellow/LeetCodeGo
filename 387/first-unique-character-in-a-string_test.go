package leetcode

import "testing"

func Test1(t *testing.T) {
	input := "leetcode"
	output := firstUniqChar(input)
	expect := 0
	if output != expect {
		t.Fail()
	}
}


func Test2(t *testing.T) {
	input := "loveleetcode"
	output := firstUniqChar(input)
	expect := 2
	if output != expect {
		t.Fail()
	}
}



func Test3(t *testing.T) {
	input := ""
	output := firstUniqChar(input)
	expect := 0
	if output != expect {
		t.Fail()
	}
}

package leetcode

import "testing"

func Test1(t *testing.T) {
	s := "aaabb"
	k := 3
	output := longestSubstring(s, k)
	expect := 3
	if output != expect {
		t.Fail()
	}
}


func Test2(t *testing.T) {
	s := "aaabb"
	k := 1
	output := longestSubstring(s, k)
	expect := 5
	if output != expect {
		t.Fail()
	}
}



func Test3(t *testing.T) {
	s := "ababacb"
	k := 3
	output := longestSubstring(s, k)
	expect := 0
	if output != expect {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	s := "aabcabb"
	k := 3
	output := longestSubstring(s, k)
	expect := 0
	if output != expect {
		t.Fail()
	}
}


func Test5(t *testing.T) {
	s := "bbaaacbd"
	k := 3
	output := longestSubstring(s, k)
	expect := 3
	if output != expect {
		t.Fail()
	}
}
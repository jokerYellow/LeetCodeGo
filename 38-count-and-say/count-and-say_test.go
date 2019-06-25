package _8_count_and_say

import "testing"

func Test1(t *testing.T) {
	input := 1
	output := countAndSay(input)
	expect := "1"
	if output != expect {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	input := 5
	output := countAndSay(input)
	expect := "111221"
	if output != expect {
		t.Fail()
	}
}

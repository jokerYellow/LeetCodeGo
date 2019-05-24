package _6_minimum_window_substring

import "testing"

func Test1(t *testing.T) {
	input := "ADOBECODEBANC"
	tt := "ABC"
	expect := "BANC"
	output := minWindow(input, tt)
	if output != expect {
		t.Fail()
	}
}


func Test2(t *testing.T) {
	input := "ANC"
	tt := "ABC"
	expect := ""
	output := minWindow(input, tt)
	if output != expect {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	input := "bbaa"
	tt := "aba"
	expect := "baa"
	output := minWindow(input, tt)
	if output != expect {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	input := "A"
	tt := "AA"
	expect := ""
	output := minWindow(input, tt)
	if output != expect {
		t.Fail()
	}
}

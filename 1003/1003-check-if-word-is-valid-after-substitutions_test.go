package _003

import "testing"

func TestNormal(t *testing.T) {
	Input := "cababc"
	Expect := false
	Output := isValid(Input)
	if Expect != Output {
		t.Fail()
	}
}

func TestRepeat(t *testing.T) {
	Input := "aaabcbcbc"
	Expect := true
	Output := isValid(Input)
	if Expect != Output {
		t.Fail()
	}
}

func TestBounds(t *testing.T) {
	Input := "a"
	Expect := false
	Output := isValid(Input)
	if Expect != Output {
		t.Fail()
	}
}
func TestBounds3(t *testing.T) {
	Input := "abc"
	Expect := true
	Output := isValid(Input)
	if Expect != Output {
		t.Fail()
	}
}

func TestBounds2(t *testing.T) {
	Input := "abcbcbc"
	Expect := false
	Output := isValid(Input)
	if Expect != Output {
		t.Fail()
	}
	Input = "abc"
	Expect = true
	Output = isValid(Input)
	if Expect != Output {
		t.Fail()
	}
}

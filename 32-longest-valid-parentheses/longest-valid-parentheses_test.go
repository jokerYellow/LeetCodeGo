package _2_longest_valid_parentheses

import "testing"

func Test1(t *testing.T) {
	if longestValidParentheses("(()") != 2 {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	if longestValidParentheses(")()())") != 4 {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	if longestValidParentheses(")()())((((()))))") != 10 {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	if longestValidParentheses("") != 0 {
		t.Fail()
	}
}

func Test5(t *testing.T) {
	if longestValidParentheses(")))(((") != 0 {
		t.Fail()
	}
}

func Test6(t *testing.T) {
	if longestValidParentheses("()(()") != 2 {
		t.Fail()
	}
}

func Test7(t *testing.T) {
	if longestValidParentheses("()()))()") != 4 {
		t.Fail()
	}
}

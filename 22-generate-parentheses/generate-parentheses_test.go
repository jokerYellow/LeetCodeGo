package _2_generate_parentheses

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

func Test1(t *testing.T) {
	input := 3
	output := generateParenthesis(input)
	expect := []string{"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()"}
	if !utils.EqualStrings(output, expect) {
		t.Fail()
	}
}

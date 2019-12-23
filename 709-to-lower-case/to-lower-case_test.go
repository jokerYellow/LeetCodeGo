package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	str    string
	expect string
}

func TestName(t *testing.T) {
	cases := []testCase{
		{"", ""},
		{"Hello", "hello"},
		{"LOVELY", "lovely"},
		{"LOVELY好", "lovely好"},
	}
	for _, item := range cases {
		output := toLowerCase(item.str)
		if output == item.expect {
			utils.Print(item.expect, output)
		}
	}
}

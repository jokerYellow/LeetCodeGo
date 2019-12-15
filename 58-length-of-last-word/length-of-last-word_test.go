package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	s      string
	expect int
}

func Test(t *testing.T) {
	cases := []testCase{
		{"Hello World", 5},
		{"World", 5},
		{"World    ", 5},
		{"", 0},
	}
	for _, c := range cases {
		output := lengthOfLastWord(c.s)
		if output != c.expect {
			utils.Print(c.expect, output)
			t.Fail()
		}
	}
}

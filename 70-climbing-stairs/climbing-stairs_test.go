package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	n      int
	expect int
}

func Test(t *testing.T) {
	cases := []testCase{
		//{45, 1},
		{2, 2},
		{3, 3},
		{4, 5},
	}
	for _, c := range cases {
		output := climbStairs(c.n)
		if output != c.expect {
			utils.Print(c.expect, output)
			t.Fail()
		}
	}
}

package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	dividend int
	divisor  int
	expect   int
}

func Test(t *testing.T) {
	cases := []testCase{
		{-2147483648, -1, 2147483647},
		{2147483648, 1, 2147483647},
		{7, -3, -2},
		{10, 3, 3},
		{9, 3, 3},
	}
	for _, item := range cases {
		output := divide(item.dividend, item.divisor)
		if output != item.expect {
			utils.Print(item.expect, output)
			t.Fail()
		}
	}
}

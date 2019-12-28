package powx_n

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	x      float64
	n      int
	expect float64
}

func Test(t *testing.T) {
	cases := []testCase{
		{2.1, 3, 9.261},
		{2, 10, 1024},
		{1, 1111111111111111110, 1},
		{2.0, -2, 0.25},
		{0, -2, 0},
		{-2, 0, 1},
	}
	for _, c := range cases {
		output := myPow(c.x, c.n)
		if output != c.expect {
			utils.Print(c.expect, output)
			t.Fail()
		}
	}
}

package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	x      int
	expect int
}

func Test(t *testing.T) {
	cases := []testCase{
		{4, 2},
		{8, 2},
		{1, 1},
		{0, 0},
	}
	for _, c := range cases {
		o := mySqrt(c.x)
		if o != c.expect {
			utils.Print(c.expect, o)
			t.Fail()
		}
	}

}

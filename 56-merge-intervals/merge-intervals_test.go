package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	intervals [][]int
	expect    [][]int
}

func Test(t *testing.T) {
	cases := []testCase{
		{
			[][]int{{1, 4}, {4, 5}, {2, 3}},
			[][]int{{1, 5},},
		},
		{
			[][]int{{1, 3}, {2, 6}, {15, 18}, {8, 10},},
			[][]int{{1, 6}, {8, 10}, {15, 18},},
		},
	}
	for _, c := range cases {
		output := merge(c.intervals)
		if !utils.CheckEqual(output, c.expect) {
			utils.Print(c.expect, output)
			t.Fail()
		}
	}
}

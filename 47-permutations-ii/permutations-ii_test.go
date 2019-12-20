package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	nums   []int
	expect [][]int
}

func Test(t *testing.T) {
	cases := []testCase{
		{[]int{1, 1, 2},
			[][]int{{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
		{[]int{},
			[][]int{},
		},
		{[]int{1, 2, 3},
			[][]int{{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{[]int{1},
			[][]int{{1},
			},
		},
	}
	for _, c := range cases {
		output := permuteUnique(c.nums)
		if !utils.CheckElementsEqual(output, c.expect) {
			utils.Print(c.expect, output)
			t.Fail()
		}
	}
}

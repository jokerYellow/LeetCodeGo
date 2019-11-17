package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	nums   []int
	target int
	expect int
}

func Test(t *testing.T) {
	cases := []testCase{
		{[]int{1, 3, 5, 6}, 4, 2},
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 6, 3},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1, 3, 5, 6}, 1110, 4},
		{[]int{1}, 1110, 1},
		{[]int{}, 6, 0},
	}
	for _, c := range cases {
		output := searchInsert(c.nums, c.target)
		if output != c.expect {
			utils.Print(c.expect, output)
			t.Fail()
		}
	}
}

package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	nums   []int
	target int
	expect []int
}

func Test(t *testing.T) {
	cases := []testCase{
		{[]int{1, 1, 1, 1, 1}, 1, []int{0, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
	}
	for _, c := range cases {
		output := searchRange(c.nums, c.target)
		if !utils.CheckEqualArr(output, c.expect) {
			utils.Print(c.expect, output)
			t.Fail()
		}
	}
}

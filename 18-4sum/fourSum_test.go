package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	nums   []int
	target int
	expect [][]int
}

func Test(t *testing.T) {
	cases := []testCase{
		//[[-3,-2,2,3],[-3,-1,1,3],[-3,0,0,3],[-3,0,1,2],[-2,-1,0,3],[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

		{[]int{1, 1, 1, 1, 1}, 4, [][]int{
			{1, 1, 1, 1},
		}},
		{[]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0, [][]int{
			{-3, -2, 2, 3},
			{-3, -1, 1, 3},
			{-3, 0, 0, 3},
			{-3, 0, 1, 2},
			{-2, -1, 0, 3},
			{-2, -1, 1, 2},
			{-2, 0, 0, 2},
			{-1, 0, 0, 1}}},
		{[]int{1, 0, -1}, 0, [][]int{}},
		{[]int{1, 0, -1, 100}, 11110, [][]int{}},
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-1, 0, 0, 1}, {-2, -1, 1, 2}, {-2, 0, 0, 2}}},
	}
	for _, item := range cases {
		output := fourSum(item.nums, item.target)
		if !utils.CheckItemsEqual(output, item.expect) {
			t.Fail()
		}
	}

}

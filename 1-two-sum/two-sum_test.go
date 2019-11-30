package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testcase struct {
	nums   []int
	target int
	expect []int
}

func Test(t *testing.T) {
	cases := []testcase{
		{[]int{2, 7, 11, 15},
			9,
			[]int{0, 1}},
		{[]int{},
			9,
			[]int{}},
	}
	for _, item := range cases {
		if !utils.CheckEqualArr(twoSum(item.nums, item.target), item.expect) {
			t.Fail()
		}
	}
}

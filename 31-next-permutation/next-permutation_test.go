package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testcase struct {
	nums   []int
	expect []int
}

func Test(t *testing.T) {
	cases := []testcase{
		{[]int{3, 3, 2, 3, 4, 5, 6, 5, 4, 3}, []int{3, 3, 2, 3, 4, 6, 3, 4, 5, 5}},
		{[]int{2, 3, 1, 3, 3}, []int{2, 3, 3, 1, 3}},
		{[]int{2, 3, 1}, []int{3, 1, 2}},
		{[]int{1, 3, 2}, []int{2, 1, 3}},
		{[]int{}, []int{}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
		{[]int{1, 1, 2}, []int{1, 2, 1}},
		{[]int{3, 3, 2, 3}, []int{3, 3, 3, 2}},
	}
	for _, item := range cases {
		output := nextPermutation(item.nums)
		if !utils.CheckEqualArr(output, item.expect) {
			t.Fail()
		}
	}
}

package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	nums1, nums2 []int
	m, n         int
	expect       []int
}

func Test(t *testing.T) {
	cases := []testCase{
		{
			make([]int, 1),
			[]int{1},
			0,
			1,
			[]int{1},
		},
		{
			[]int{1},
			[]int{},
			1,
			0,
			[]int{1},
		},
		{
			[]int{},
			[]int{},
			0,
			0,
			[]int{},
		},
		{
			[]int{1, 2, 3, 0, 0, 0},
			[]int{2, 5, 6},
			3,
			3,
			[]int{1, 2, 2, 3, 5, 6},
		},
	}
	for _, item := range cases {
		merge(item.nums1, item.m, item.nums2, item.n)
		if !utils.CheckEqualArr(item.nums1, item.expect) {
			utils.Print(item.expect, item.nums1)
			t.Fail()
		}
	}
}

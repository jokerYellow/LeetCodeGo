package leetcode

import "testing"

type testCase struct {
	nums   []int
	expect int
}

func Test(t *testing.T) {
	cases := []testCase{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1, 1, 1}, 3},
		{[]int{-3, -1, -1}, -1},
		{[]int{}, 0},
	}
	for _, item := range cases {
		output := maxSubArray(item.nums)
		if output != item.expect {
			t.Fail()
		}
	}
}

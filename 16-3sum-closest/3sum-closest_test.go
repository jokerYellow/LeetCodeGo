package leetcode

import "testing"

type testcase struct {
	nums   []int
	target int
	expect int
}

func Test(t *testing.T) {
	cases := []testcase{
		{[]int{0, 2, 1, -3}, 1, 0},
		{[]int{-1, 2, 1, -4}, 1, 2},
		{[]int{-1, 2, 1, -4}, 0, -1},
	}
	for _, c := range cases {
		output := threeSumClosest(c.nums, c.target)
		if output != c.expect {
			t.Fail()
		}
	}
}

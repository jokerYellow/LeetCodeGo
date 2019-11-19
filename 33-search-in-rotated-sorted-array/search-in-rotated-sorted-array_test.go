package leetcode

import "testing"

type testCase struct {
	nums   []int
	target int
	expect int
}

func Test(t *testing.T) {
	cases := []testCase{
		{[]int{3, 1}, 1, 1},
		{[]int{4, 5, 6, 7, 8, 1, 2, 3}, 8, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{}, 0, -1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1}}
	for _, c := range cases {
		output := search(c.nums, c.target)
		if output != c.expect {
			t.Fail()
		}
	}
}

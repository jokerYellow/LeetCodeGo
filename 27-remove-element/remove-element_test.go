package leetcode

import "testing"

type testCase struct {
	nums   []int
	val    int
	expect int
}

func Test(t *testing.T) {
	cases := []testCase{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}
	for _, item := range cases {
		output := removeElement(item.nums, item.val)
		if output != item.expect {
			t.Fail()
		}
	}
}

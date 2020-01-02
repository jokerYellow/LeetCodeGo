package leetcode

import "testing"

type testCase struct {
	nums   []int
	expect bool
}

func Test(t *testing.T) {
	cases := []testCase{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	for _, c := range cases {
		output := containsDuplicate(c.nums)
		if output != c.expect {
			t.Fail()
		}
	}
}

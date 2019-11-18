package leetcode

import "testing"

type testcase struct {
	nums   []int
	expect bool
}

func Test(t *testing.T) {
	cases := []testcase{
		testcase{[]int{3, 2, 1, 0, 4}, false},
		testcase{[]int{2, 0, 0}, true},
		testcase{[]int{1, 2, 3}, true},
		testcase{[]int{0}, true},
		testcase{[]int{0, 1}, false},
		testcase{[]int{100, 0, 0, 1, 0, 0, 1}, true},
	}
	for _, item := range cases {
		output := canJump(item.nums)
		if output != item.expect {
			t.Fail()
		}
	}
}

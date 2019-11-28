package leetcode

import "testing"

type testCase struct {
	n      int
	expect float64
}

func Test(t *testing.T) {
	cases := []testCase{
		{1, 1},
		{2, 0.5},
		{3, 0.5},
		{4, 0.5},
	}
	for _, item := range cases {
		output := nthPersonGetsNthSeat(item.n)
		if output != item.expect {
			t.Fail()
		}
	}
}

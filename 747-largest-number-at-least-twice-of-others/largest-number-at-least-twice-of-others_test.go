package leetcode

import "testing"

type testCase struct {
	nums   []int
	expect int
}

func Test(t *testing.T) {
	testCases := []testCase{
		{[]int{3, 6, 1, 0}, 1},
		{[]int{1, 2, 3, 4}, -1},
	}
	for _, c := range testCases {
		output := dominantIndex(c.nums)
		if output != c.expect {
			t.Fail()
		}
	}
}

package leetcode

import (
	"fmt"
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	nums   []int
	k      int
	expect bool
}

func Test(t *testing.T) {
	cases := []testCase{
		{[]int{99, 99}, 2, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{}, 1, false},
	}
	for _, c := range cases {
		output := containsNearbyDuplicate(c.nums, c.k)
		if output != c.expect {
			fmt.Println(c.nums)
			utils.Print(c.expect, output)
			t.Fail()
		}
	}
}

package leetcode

import (
	"fmt"
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	nums   []int
	k      int
	t      int
	expect bool
}

func Test(t *testing.T) {
	cases := []testCase{
		{[]int{1, 5, 9, 1, 5, 9}, 2, 3, false},
		{[]int{1, 0, 1, 1}, 1, 2, true},
		{[]int{1, 2, 3, 1}, 3, 0, true},
	}
	for _, c := range cases {
		output := containsNearbyAlmostDuplicate(c.nums, c.k, c.t)
		if output != c.expect {
			fmt.Println(c.nums)
			utils.Print(c.expect, output)
			t.Fail()
		}
	}
}

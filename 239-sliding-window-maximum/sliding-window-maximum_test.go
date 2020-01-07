package leetcode

import (
	"fmt"
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	nums   []int
	k      int
	expect []int
}

func Test(t *testing.T) {
	cases := []testCase{
		{[]int{1, 3, 1, 2, 0, 5}, 3, []int{3, 3, 2, 5}},
		{[]int{1, -1}, 1, []int{1, -1}},
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
	}
	for _, c := range cases {
		o := maxSlidingWindow(c.nums, c.k)
		if !utils.CheckEqualArr(o, c.expect) {
			fmt.Println(c.nums, c.k)
			utils.Print(c.expect, o)
			t.Fail()
		}
	}
}

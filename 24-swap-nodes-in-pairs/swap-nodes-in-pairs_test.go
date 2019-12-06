package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	input  []int
	expect []int
}

func Test(t *testing.T) {
	cases := []testCase{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{2, 1, 3}},
		{[]int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
	}
	for _, item := range cases {
		output := swapPairs(utils.GenerateLinkList(item.input))
		if !utils.CheckEqualLink(utils.GenerateLinkList(item.expect), output) {
			t.Fail()
		}
	}
}

package leetcode

import (
	"testing"

	"github.com/jokerYellow/leetcode/utils"
)

type testCase struct {
	links  []int
	n      int
	expect []int
}

func Test(t *testing.T) {
	cases := []testCase{
		{[]int{1, 2}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{2}},
		{[]int{1}, 2, []int{1}},
		{[]int{}, 2, []int{}},
	}
	for _, c := range cases {
		head := utils.GenerateLinkList(c.links)
		head = removeNthFromEnd(head, c.n)
		if !utils.CheckEqualLink(head, utils.GenerateLinkList(c.expect)) {
			t.Fail()
		}
	}
}


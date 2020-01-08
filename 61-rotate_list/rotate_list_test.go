package rotate_list

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	head   *utils.ListNode
	k      int
	expect *utils.ListNode
}

func Test(t *testing.T) {
	cases := []testCase{
		{utils.GenerateLinkList([]int{1, 2, 3, 4, 5}), 2, utils.GenerateLinkList([]int{4, 5, 1, 2, 3})},
		{utils.GenerateLinkList([]int{0, 1, 2}), 4, utils.GenerateLinkList([]int{2, 0, 1})},
		{utils.GenerateLinkList([]int{0, 1, 2}), 0, utils.GenerateLinkList([]int{0, 1, 2})},
		{utils.GenerateLinkList([]int{1}), 199, utils.GenerateLinkList([]int{1})},
		{utils.GenerateLinkList([]int{}), 199, utils.GenerateLinkList([]int{})},
	}
	for _, c := range cases {
		o := rotateRight(c.head, c.k)
		if !utils.CheckEqualLink(o, c.expect) {
			utils.Print(c.expect, o)
			t.Fail()
		}
	}
}

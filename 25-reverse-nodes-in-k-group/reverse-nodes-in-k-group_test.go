package leetcode

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
		{utils.GenerateLinkList([]int{}), 2, utils.GenerateLinkList([]int{})},
		{utils.GenerateLinkList([]int{1}), 2, utils.GenerateLinkList([]int{1})},
		{utils.GenerateLinkList([]int{1, 2, 3, 4, 5}), 2, utils.GenerateLinkList([]int{2, 1, 4, 3, 5})},
		{utils.GenerateLinkList([]int{1, 2, 3, 4, 5}), 3, utils.GenerateLinkList([]int{3, 2, 1, 4, 5})},
	}
	for _, item := range cases {
		output := reverseKGroup(item.head, item.k)
		if !utils.CheckEqualLink(output, item.expect) {
			utils.Print(item.expect, output)
			t.Fail()
		}
	}
}

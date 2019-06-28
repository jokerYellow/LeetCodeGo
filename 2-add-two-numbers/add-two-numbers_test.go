package __add_two_numbers

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

func Test1(t *testing.T) {
	l1 := utils.GenerateLinkList([]int{2, 4, 3, 1})
	l2 := utils.GenerateLinkList([]int{5, 6, 4})
	l3 := addTwoNumbers(l1, l2)
	if !utils.CheckEqualLink(l3, utils.GenerateLinkList([]int{7, 0, 8, 1})) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	l1 := utils.GenerateLinkList([]int{2, 4, 3, 1})
	l2 := utils.GenerateLinkList([]int{})
	l3 := addTwoNumbers(l1, l2)
	if !utils.CheckEqualLink(l3, utils.GenerateLinkList([]int{2, 4, 3, 1})) {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	l1 := utils.GenerateLinkList([]int{})
	l2 := utils.GenerateLinkList([]int{})
	l3 := addTwoNumbers(l1, l2)
	if !utils.CheckEqualLink(l3, utils.GenerateLinkList([]int{})) {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	l1 := utils.GenerateLinkList([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	l2 := utils.GenerateLinkList([]int{9, 9, 9, 9, 9, 9})
	l3 := addTwoNumbers(l1, l2)
	if !utils.CheckEqualLink(l3, utils.GenerateLinkList([]int{0, 1, 1, 1, 1, 1, 2, 1, 1, 1})) {
		t.Fail()
	}
}

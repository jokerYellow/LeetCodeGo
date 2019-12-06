package leetcode

import "github.com/jokerYellow/leetcode/utils"

func swapPairs(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, first *utils.ListNode
	current := head
	newHead := head.Next
	for current != nil {
		if first == nil {
			first = current
			current = current.Next
		} else {
			first.Next = current.Next
			current.Next = first
			if pre != nil {
				pre.Next = current
			}
			pre = first
			current = first.Next
			first = nil
		}
	}
	return newHead
}

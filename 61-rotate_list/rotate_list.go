package rotate_list

import (
	"github.com/jokerYellow/leetcode/utils"
)

func rotateRight(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil {
		return nil
	}
	var nodes []*utils.ListNode
	c := head
	for c != nil {
		nodes = append(nodes, c)
		c = c.Next
	}
	length := len(nodes)
	rotateCount := k % length
	c = head
	for i := len(nodes) - 1; i >= (len(nodes) - rotateCount); i-- {
		n := nodes[i]
		n.Next = c
		if i-1 >= 0 {
			nodes[i-1].Next = nil
		}
		c = n
	}
	return c
}

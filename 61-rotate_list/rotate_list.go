package rotate_list

import (
	"github.com/jokerYellow/leetcode/utils"
)

//circle the link,time:O(n),space:O(1)
func rotateRight(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil || k == 0 {
		return head
	}
	c := head
	length := 1
	for c.Next != nil {
		c = c.Next
		length++
	}
	c.Next = head
	rotateCount := k % length
	for i := 0; i < length - rotateCount; i++ {
		c = c.Next
	}
	t := c.Next
	c.Next = nil
	return t
}

//hold the nodes,time:O(n),space:O(n)
func _rotateRight(head *utils.ListNode, k int) *utils.ListNode {
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

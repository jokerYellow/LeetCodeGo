package leetcode
/*
https://leetcode.com/problems/reverse-nodes-in-k-group/
25. Reverse Nodes in k-Group
Hard

1573

309

Add to List

Share
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

Example:

Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5

Note:

Only constant extra memory is allowed.
You may not alter the values in the list's nodes, only nodes itself may be changed.
 */
import (
	"github.com/jokerYellow/leetcode/utils"
)
//todo: understand and rewrite
func reverseKGroup(head *utils.ListNode, k int) *utils.ListNode {
	curr := head
	count := 0
	for curr != nil && count != k {
		curr = curr.Next
		count++
	}
	if count == k {
		curr = reverseKGroup(curr, k)
		for count > 0 {
			tmp := head.Next
			head.Next = curr
			curr = head
			head = tmp
			count--
		}
		head = curr
	}
	return head
}

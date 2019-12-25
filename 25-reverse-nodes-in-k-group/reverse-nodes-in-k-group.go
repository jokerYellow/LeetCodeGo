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

func reverseKGroup(head *utils.ListNode, k int) *utils.ListNode {
	count := 0
	var preHead = head
	for count < k-1 && preHead != nil  {
		preHead = preHead.Next
		count++
	}
	if count != k-1 || preHead == nil {
		return head
	}
	preHead.Next = reverseKGroup(preHead.Next, k)
	for count > 0 {
		next := head.Next
		head.Next = preHead.Next
		preHead.Next = head
		head = next
		count--
	}
	return head
}

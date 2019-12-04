package leetcode

import "github.com/jokerYellow/leetcode/utils"

/*
https://leetcode.com/problems/remove-nth-node-from-end-of-list/
19. Remove Nth Node From End of List
Medium

2348

177

Favorite

Share
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?
*/

func removeNthFromEnd(head *utils.ListNode, n int) *utils.ListNode {
	if n == 0 {
		return head
	}

	//first find normal index to be removed
	length := 0
	current := head
	for current != nil {
		current = current.Next
		length++
	}
	count := length - n
	if count < 0 {
		return head
	}
	//remove head
	if count == 0 {
		result := head.Next
		head.Next = nil
		return result
	}
	//remove normal index
	current = head
	pre := head
	for count > 0 {
		count--
		pre = current
		current = current.Next
	}
	next := current.Next
	current.Next = nil
	if pre != nil {
		pre.Next = next
	}
	return head
}

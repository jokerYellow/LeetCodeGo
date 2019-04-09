package leetcode

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
82. Remove Duplicates from Sorted List II
Medium

748

70

Favorite

Share
Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list.

Example 1:

Input: 1->2->3->3->4->4->5
Output: 1->2->5
Example 2:

Input: 1->1->1->2->3
Output: 2->3
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var newHead *ListNode
	var last *ListNode
	for head != nil {
		if head.Next != nil && head.Next.Val == head.Val {
			for head != nil {
				if head.Next != nil && head.Next.Val != head.Val {
					head = head.Next
					break
				}
				head = head.Next
			}
		} else {
			t := new(ListNode)
			t.Val = head.Val
			if newHead == nil {
				newHead = t
				last = t
			} else {
				last.Next = t
				last = last.Next
			}
			head = head.Next
		}
	}
	return newHead
}

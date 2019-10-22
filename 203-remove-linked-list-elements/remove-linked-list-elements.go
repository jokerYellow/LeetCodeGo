package leetcode

/*
https://leetcode.com/problems/remove-linked-list-elements/
Easy

760

43

Favorite

Share
Remove all elements from a linked list of integers that have value val.

Example:

Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	return removeElements2(head, val)
}
func removeElements1(head *ListNode, val int) *ListNode {
	var newHead *ListNode
	var last *ListNode
	for head != nil {
		if head.Val != val {
			t := &ListNode{Val: head.Val}
			if newHead == nil {
				newHead = t
				last = t
			} else {
				last.Next = t
				last = last.Next
			}
		}
		head = head.Next
	}
	return newHead
}

func removeElements2(head *ListNode, val int) *ListNode {
	virHead := &ListNode{Next: head}
	first := virHead
	for first != nil {
		for first.Next != nil && first.Next.Val == val {
			first.Next = first.Next.Next
		}
		first = first.Next
	}
	return virHead.Next
}

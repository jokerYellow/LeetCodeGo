package leetcode

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-list/

83. Remove Duplicates from Sorted List
Easy

705

77

Favorite

Share
Given a sorted linked list, delete all duplicates such that each element appear only once.

Example 1:

Input: 1->1->2
Output: 1->2
Example 2:

Input: 1->1->2->3->3
Output: 1->2->3
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil{
		return head
	}
	var current *ListNode
	var move = head
	for {
		if current == nil {
			current = move
		} else if move.Val != current.Val {
			current.Next = move
			current = current.Next
		}
		move = move.Next
		if move == nil {
			break
		}
		current.Next = nil
	}
	return head
}

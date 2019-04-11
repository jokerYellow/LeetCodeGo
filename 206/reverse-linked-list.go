package leetcode

/*
https://leetcode.com/problems/reverse-linked-list/
206. Reverse Linked List
Easy

2144

60

Favorite

Share
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	return reverseList2(head)
}

//recursively
func reverseList1(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		return reverse(head,head.Next)
	}else{
		return head
	}
}

func reverse(head *ListNode,next *ListNode)*ListNode  {
	var rt *ListNode
	if next.Next == nil {
		rt = next
	}else{
		rt = reverse(next,next.Next)
	}
	next.Next = head
	head.Next = nil
	return rt
}

//iteratively
func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	newHead := head
	head = head.Next
	newHead.Next = nil
	for head != nil {
		var next *ListNode
		if head.Next != nil {
			next = head.Next
		}
		head.Next = newHead
		newHead = head
		if next != nil{
			head = next
		}else{
			break
		}
	}
	return newHead
}

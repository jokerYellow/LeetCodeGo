package leetcode

/*
https://leetcode.com/problems/palindrome-linked-list/
234. Palindrome Linked List
Easy

1497

226

Favorite

Share
Given a singly linked list, determine if it is a palindrome.

Example 1:

Input: 1->2
Output: false
Example 2:

Input: 1->2->2->1
Output: true
Follow up:
Could you do it in O(n) time and O(1) space?

Accepted
242,480
Submissions
681,513
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	if head.Next == nil {
		return true
	}
	slow := head
	fast := head
	var arr []int
	for fast != nil {
		if fast.Next == nil {
			fast = nil
			slow = slow.Next
			break
		} else if fast.Next.Next == nil {
			fast = fast.Next
			arr = append(arr, slow.Val)
			slow = slow.Next
			break
		} else {
			fast = fast.Next.Next
		}
		arr = append(arr, slow.Val)
		slow = slow.Next
	}
	middle := slow

	length := len(arr) - 1

	for {
		if middle == nil {
			break
		}
		if middle.Val != arr[length] {
			return false
		}
		middle = middle.Next
		length--
	}

	return true
}

package leetcode

/*
https://leetcode.com/problems/merge-two-sorted-lists/

21. Merge Two Sorted Lists
Easy

2059

284

Favorite

Share
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
Accepted
541,003
Submissions
1,165,640
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	h1 := l1
	h2 := l2
	var lastNode *ListNode
	var head *ListNode
	for h1 != nil || h2 != nil {
		appendNode := new(ListNode)
		if h1 == nil {
			appendNode.Val = h2.Val
			h2 = h2.Next
		} else if h2 == nil {
			appendNode.Val = h1.Val
			h1 = h1.Next
		} else if h1.Val < h2.Val {
			appendNode.Val = h1.Val
			h1 = h1.Next
		} else if h1.Val >= h2.Val {
			appendNode.Val = h2.Val
			h2 = h2.Next
		}
		if lastNode == nil {
			lastNode = appendNode
			head = appendNode
		} else {
			lastNode.Next = appendNode
			lastNode = lastNode.Next
		}
	}
	return head
}

package leetcode

/*
https://leetcode.com/problems/merge-k-sorted-lists/
23. Merge k Sorted Lists
Hard

2248

143

Favorite

Share
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var rt *ListNode
	for _, v := range lists {
		if rt == nil {
			rt = v
		} else {
			rt = mergeTwoLists(rt, v)
		}
	}
	return rt
}

func mergeTwoLists(head1, head2 *ListNode) *ListNode {
	var current *ListNode
	var head *ListNode
	for head1 != nil || head2 != nil {
		var t *ListNode
		if head1 == nil {
			t = head2
			head2 = head2.Next
		} else if head2 == nil {
			t = head1
			head1 = head1.Next
		} else {
			if head1.Val >= head2.Val {
				t = head2
				head2 = head2.Next
			} else {
				t = head1
				head1 = head1.Next
			}
		}
		if current == nil {
			current = t
			head = current
		} else {
			current.Next = t
			current = current.Next
		}
	}
	return head
}

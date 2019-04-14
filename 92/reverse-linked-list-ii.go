package leetcode

/*
https://leetcode.com/problems/reverse-linked-list-ii/
92. Reverse Linked List II
Medium

1086

80

Favorite

Share
Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
1,2,3,4,5
1,3,2,4,5

c = 4
4.next = 5
4.next = 1.next 3
1.next = 4


1,4,3,2,5

*/
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	var dummy = &ListNode{Next: head}
	var mNode *ListNode
	var newHead *ListNode
	index := 0
	c := dummy
	for c != nil {
		if index == m-1 {
			newHead = c
			mNode = c.Next
			c = c.Next
		} else if index > m && index <= n {
			next := c.Next
			c.Next = newHead.Next
			newHead.Next = c
			if index == n {
				mNode.Next = next
			}
			c = next
		} else {
			c = c.Next
		}
		index += 1
	}
	return dummy.Next
}

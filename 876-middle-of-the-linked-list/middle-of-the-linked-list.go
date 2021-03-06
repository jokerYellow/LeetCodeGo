package leetcode

/*
https://leetcode.com/problems/middle-of-the-linked-list/
876. Middle of the Linked List
Easy

433

33

Favorite

Share
Given a non-empty, singly linked list with head node head, return a middle node of linked list.

If there are two middle nodes, return the second middle node.



Example 1:

Input: [1,2,3,4,5]
Output: Node 3 from this list (Serialization: [3,4,5])
The returned node has value 3.  (The judge's serialization of this node is [3,4,5]).
Note that we returned a ListNode object ans, such that:
ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, and ans.next.next.next = NULL.
Example 2:

Input: [1,2,3,4,5,6]
Output: Node 4 from this list (Serialization: [4,5,6])
Since the list has two middle nodes with values 3 and 4, we return the second one.


Note:

The number of nodes in the given list will be between 1 and 100.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	return middleNode2(head)
}

func middleNode1(head *ListNode) *ListNode {
	var length = 0
	c := head
	for {
		if c == nil {
			break
		}
		length++
		c = c.Next
	}
	index := length / 2
	rt := head
	current := 0
	for {
		if index == current {
			break
		}
		current++
		rt = rt.Next
	}
	return rt
}

func middleNode2(head *ListNode) *ListNode {
	p1 := head
	p2 := head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1
}

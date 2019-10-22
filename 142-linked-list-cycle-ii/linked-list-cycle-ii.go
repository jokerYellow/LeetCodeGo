package leetcode

/*
https://leetcode.com/problems/linked-list-cycle-ii/

142. Linked List Cycle II
Medium

1270

87

Favorite

Share
Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

Note: Do not modify the linked list.



Example 1:

Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.


Example 2:

Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.


Example 3:

Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.




Follow up:
Can you solve it without using extra space?
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	head1 := hasCycle(head)
	if head1 == nil {
		return nil
	}
	if head1 == head {
		return head1
	}
	rt := findIntersection(head1, head, head1)
	return rt
}

func hasCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p1 := head
	p2 := head
	for {
		if p2.Next != nil {
			p2 = p2.Next
		} else {
			return nil
		}
		if p2.Next != nil {
			p2 = p2.Next
		} else {
			return nil
		}
		p1 = p1.Next
		if p1 == p2 {
			return p1
		}
	}
}

func findIntersection(head1, head2, end *ListNode) *ListNode {
	p1 := head1
	p2 := head2
	for {
		if p1 == p2 {
			return p1
		}
		if p1.Next == p2.Next {
			return p1.Next
		}
		if p1.Next == end {
			p1 = head2
		} else {
			p1 = p1.Next
		}

		p2 = p2.Next
	}
}

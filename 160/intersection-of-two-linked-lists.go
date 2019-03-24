package leetcode

/*
https://leetcode.com/problems/intersection-of-two-linked-lists

160. Intersection of Two Linked Lists
Easy

1913

148

Favorite

Share
Write a program to find the node at which the intersection of two singly linked lists begins.

For example, the following two linked lists:


begin to intersect at node c1.



Example 1:


Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
Output: Reference of the node with value = 8
Input Explanation: The intersected node's value is 8 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [4,1,8,4,5]. From the head of B, it reads as [5,0,1,8,4,5]. There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.


Example 2:


Input: intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
Output: Reference of the node with value = 2
Input Explanation: The intersected node's value is 2 (note that this must not be 0 if the two lists intersect). From the head of A, it reads as [0,9,1,2,4]. From the head of B, it reads as [3,2,4]. There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.


Example 3:


Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
Output: null
Input Explanation: From the head of A, it reads as [2,6,4]. From the head of B, it reads as [1,5]. Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.
Explanation: The two lists do not intersect, so return null.


Notes:

If the two linked lists have no intersection at all, return null.
The linked lists must retain their original structure after the function returns.
You may assume there are no cycles anywhere in the entire linked structure.
Your code should preferably run in O(n) time and use only O(1) memory.
*/

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	return getIntersectionNode2(headA, headB)
}

/*
 2,6,3,4,5
 1,2,4,5

 2,6,3,4,1,2,4,5
 1,2,4,2,6,3,4,5
*/
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	ha := headA
	hb := headB
	for ha != hb {
		if ha == nil {
			ha = headB
		} else {
			ha = ha.Next
		}

		if hb == nil {
			hb = headA
		} else {
			hb = hb.Next
		}
	}
	return ha
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	lengthA := length(headA)
	lengthB := length(headB)

	var nodeA *ListNode
	var nodeB *ListNode
	if lengthA > lengthB {
		nodeA = move(headA, lengthA-lengthB)
		nodeB = headB
	} else {
		nodeB = move(headB, -lengthA+lengthB)
		nodeA = headA
	}

	for {
		if nodeA == nodeB {
			break
		} else {
			nodeA = move(nodeA, 1)
			nodeB = move(nodeB, 1)
			if nodeA == nil {
				break
			}
		}
	}
	return nodeA
}

func move(node *ListNode, step int) *ListNode {
	if step == 0 {
		return node
	}
	if node == nil {
		return nil
	}
	var rt = node
	for {
		if rt.Next == nil {
			return nil
		}
		rt = rt.Next
		step--
		if step == 0 {
			break
		}
	}
	return rt
}

func length(head *ListNode) int {
	if head == nil {
		return 0
	}
	length := 1
	for {
		if head.Next != nil {
			length++
			head = head.Next
		} else {
			break
		}
	}
	return length
}

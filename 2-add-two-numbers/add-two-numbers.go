package __add_two_numbers

import (
	"github.com/jokerYellow/leetcode/utils"
)

/*
https://leetcode.com/problems/add-two-numbers/
2. Add Two Numbers
Medium

5372

1381

Favorite

Share
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

*/

func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	head := &utils.ListNode{}
	currentNode := head
	count := 0
	for {
		if l1 == nil && l2 == nil {
			if count > 0 {
				currentNode.Next = &utils.ListNode{Val: count}
				currentNode = currentNode.Next
			}
			break
		}
		if l1 != nil {
			count += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			count += l2.Val
			l2 = l2.Next
		}
		value := count % 10
		count = count / 10
		currentNode.Next = &utils.ListNode{Val: value}
		currentNode = currentNode.Next
	}
	return head.Next
}

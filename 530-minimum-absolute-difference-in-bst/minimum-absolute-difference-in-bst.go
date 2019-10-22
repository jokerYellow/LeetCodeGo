package leetCode

/*
530. Minimum Absolute Difference in BST
Easy

474

35

Favorite

Share
Given a binary search tree with non-negative values, find the minimum absolute difference between values of any two nodes.

Example:

Input:

   1
    \
     3
    /
   2

Output:
1

Explanation:
The minimum absolute difference is 1, which is the difference between 2 and 1 (or between 2 and 3).


Note: There are at least two nodes in this BST.
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func absolute(t int) int {
	if t < 0 {
		return -t
	}
	return t
}

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min := 0
	current := -1
	var midLoopTree func(t *TreeNode)
	midLoopTree = func(t *TreeNode) {
		if t.Left != nil {
			midLoopTree(t.Left)
		}
		if current == -1 {
			current = t.Val
		} else {
			d := absolute(current - t.Val)
			if min == 0 {
				min = d
			} else if min > d {
				min = d
			}
			current = t.Val
		}

		if t.Right != nil {
			midLoopTree(t.Right)
		}
	}
	midLoopTree(root)
	return min
}

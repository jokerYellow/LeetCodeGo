package _4_binary_tree_inorder_traversal

/*
https://leetcode.com/problems/binary-tree-inorder-traversal/
94. Binary Tree Inorder Traversal
Medium

1638

71

Favorite

Share
Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return inOrder(root)
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var rt []int
	if root.Left != nil {
		t := inOrder(root.Left)
		rt = append(rt, t...)
	}
	rt = append(rt, root.Val)
	if root.Right != nil {
		t := inOrder(root.Right)
		rt = append(rt, t...)
	}
	return rt
}

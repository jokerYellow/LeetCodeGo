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

type stack struct {
	items []*TreeNode
}

func (s *stack) pop() *TreeNode {
	top := s.top()
	if top != nil {
		s.items = s.items[:len(s.items)-1]
	}
	return top
}

func (s *stack) top() *TreeNode {
	if len(s.items) == 0 {
		return nil
	}
	return s.items[len(s.items)-1]
}

func (s *stack) push(item *TreeNode) {
	s.items = append(s.items, item)
}

func (s *stack) isEmpty() bool {
	return len(s.items) == 0
}

//Iteratively
func inorderTraversal(root *TreeNode) []int {
	s := stack{}
	var sort []int
	current := root
	for current != nil || !s.isEmpty() {
		if current != nil {
			s.push(current)
			for current.Left != nil {
				s.push(current.Left)
				current = current.Left
			}
		}
		for !s.isEmpty() {
			top := s.pop()
			if top != nil {
				sort = append(sort, top.Val)
				current = top.Right
				break
			}
		}

	}
	return sort
}

//Recursive
func inorderTraversal1(root *TreeNode) []int {
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

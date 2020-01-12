package balanced_binary_tree

/*
110. Balanced Binary Tree
Easy

1667

144

Add to List

Share
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

a binary tree in which the left and right subtrees of every node differ in height by no more than 1.



Example 1:

Given the following tree [3,9,20,null,null,15,7]:

    3
   / \
  9  20
    /  \
   15   7
Return true.

Example 2:

Given the following tree [1,2,2,3,3,null,null,4,4]:

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
Return false.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lh := height(root.Left)
	lr := height(root.Right)
	if abs(lh-lr) <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	} else {
		return false
	}
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(height(root.Left), height(root.Right))
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

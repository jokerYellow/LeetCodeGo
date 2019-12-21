package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}
	val := sum - root.Val
	if root.Left != nil && hasPathSum(root.Left, val) {
		return true
	}
	if root.Right != nil && hasPathSum(root.Right, val) {
		return true
	}
	return false
}

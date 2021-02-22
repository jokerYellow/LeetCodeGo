package leetcode

import (
	utils "github.com/jokerYellow/leetcode/utils"
)

func minDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	var less int
	if root.Left != nil && root.Right == nil {
		less = minDepth(root.Left)
	} else if root.Left == nil && root.Right != nil {
		less = minDepth(root.Right)
	} else {
		less = min(minDepth(root.Left), minDepth(root.Right))
	}
	return less + 1
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

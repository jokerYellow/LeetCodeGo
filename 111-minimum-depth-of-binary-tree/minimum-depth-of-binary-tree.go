package leetcode

import (
	utils "github.com/jokerYellow/leetcode/utils"
)

func minDepth(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right == nil {
		return minDepth(root.Left) + 1
	} else if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	} else {
		leftDepth := minDepth(root.Left)
		rightDepth := minDepth(root.Right)
		var less int
		if leftDepth < rightDepth {
			less = leftDepth
		} else {
			less = rightDepth
		}
		return less + 1
	}
}

package univalued-binary-tree

type TreeNode struct{
	Val int 
	Left *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode)bool{
	return preTree(root,root.Val)	
}

func preTree(root *TreeNode,v int)bool{
	if root == nil {
		return true
	}
	if root.Val != v {
		return false
	}
	return preTree(root.Left,v) && preTree(root.Right,v)
}
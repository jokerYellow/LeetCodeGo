package _4_binary_tree_inorder_traversal

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	root := &TreeNode{Val: 1}

	t2 := &TreeNode{Val: 2}
	t3 := &TreeNode{Val: 3}
	root.Right = t2
	t2.Left = t3

	expect := []int{1, 3, 2}
	output := inorderTraversal(root)
	fmt.Println(expect, output)
	if !equal(expect, output) {
		t.Fail()
	}
}

func equal(l, r []int) bool {
	if len(l) != len(r) {
		return false
	}
	for i, v := range l {
		if v != r[i] {
			return false
		}
	}
	return true
}

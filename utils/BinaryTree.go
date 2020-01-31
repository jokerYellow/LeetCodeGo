package utils

import "math"

const null string = "null"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func (left *TreeNode) Equal(right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else {
		return left.Val == right.Val && left.Left.Equal(right.Left) && left.Right.Equal(right.Right)
	}
}

/*
generate tree from int or null data
*/
func NewTree(vals []interface{}) *TreeNode {
	root := new(TreeNode)
	if len(vals) == 0 {
		return root
	}
	setValue(root, 0, vals)
	return root
}

func setValue(root *TreeNode, index int, vals []interface{}) {
	if len(vals) <= index {
		return
	}
	if root == nil {
		return
	}
	if v, ok := vals[index].(int); ok {
		root.Val = v
		leftIndex := index*2 + 1
		if leftIndex < len(vals) {
			if _, ok := vals[leftIndex].(int); ok {
				root.Left = new(TreeNode)
				setValue(root.Left, leftIndex, vals)
			}
		}
		rightIndex := index*2 + 2
		if rightIndex < len(vals) {
			root.Right = new(TreeNode)
			if _, ok := vals[rightIndex].(int); ok {
				setValue(root.Right, rightIndex, vals)
			}
		}
	}
}

/*
		i1
	i2		  i3
i4	   i5 i6	  i7

		i1
	i2
i4	   i5
1,2,null,4,5,null,null
1,2,null,4,5

*/
func Description(root *TreeNode) []interface{} {
	d := depth(root)
	vals := make([]interface{}, int(math.Pow(float64(2), float64(d)))-1)
	setValues(0, root, vals)
	index := 0
	for i, c := range vals {
		if c == nil {
			vals[i] = "null"
		} else {
			index = i
		}
	}
	return vals[:index+1]
}

func setValues(index int, r *TreeNode, vals []interface{}) {
	if r == nil {
		return
	}
	vals[index] = r.Val
	setValues(index*2+1, r.Left, vals)
	setValues(index*2+2, r.Right, vals)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(depth(root.Left), depth(root.Right))
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

func checkEqualTreeValsArray(l, r []interface{}) bool {
	for i, v := range l {
		if len(r) <= i {
			return false
		}
		switch v := v.(type) {
		case int:
			if vr, ok := r[i].(int); ok {
				if vr != v {
					return false
				}
			} else {
				return false
			}
		case string:
			if vr, ok := r[i].(string); ok {
				if vr != v {
					return false
				}
			} else {
				return false
			}
		default:
			return false
		}
	}
	return true
}

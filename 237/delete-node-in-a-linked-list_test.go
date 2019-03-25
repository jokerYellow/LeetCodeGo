package leetcode

import (
	"testing"
)

func TestNormal(t *testing.T) {
	n1 := new(ListNode)
	n1.Var = 2

	n2 := new(ListNode)
	n2.Var = 3

	n3 := new(ListNode)
	n3.Var = 4
	n1.Next = n2
	n2.Next = n3

	deleteNode(n2)

	if listArrCheck(n1, []*ListNode{n1, n3}) == false {
		t.Fail()
	}
}

func listArrCheck(n *ListNode, arr []*ListNode) bool {
	index := 0
	for {
		if n == nil {
			break
		}
		if n.Next == nil {
			break
		}
		if n == arr[index] {
			n = n.Next
		} else {
			return false
		}
	}
	return true
}

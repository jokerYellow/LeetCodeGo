package leetcode

import "testing"

func Test1(t *testing.T) {
	input := GenerateLinkList([]int{1, 2, 3, 4, 5})
	delete := 5
	expect := GenerateLinkList([]int{1, 2, 3, 4})
	output := removeElements(input, delete)
	if !checkEqualLink(expect, output) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	delete := 5
	output := removeElements(nil, delete)
	if !checkEqualLink(output, nil) {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	input := GenerateLinkList([]int{1, 1, 1, 1, 1})
	delete := 1
	expect := GenerateLinkList([]int{})
	output := removeElements(input, delete)
	if !checkEqualLink(expect, output) {
		t.Fail()
	}
}

func checkEqualLink(l1, l2 *ListNode) bool {
	h1 := l1
	h2 := l2
	if l1 == nil && l2 != nil {
		return false
	} else if l1 != nil && l2 == nil {
		return false
	}
	for h1 != nil && h2 != nil {
		if h1.Val != h2.Val {
			return false
		}
		h1 = h1.Next
		h2 = h2.Next
	}
	return true
}

func GenerateLinkList(arr []int) *ListNode {
	var link *ListNode
	var head *ListNode
	for v := range arr {
		l := new(ListNode)
		l.Val = arr[v]
		if link == nil {
			link = l
			head = link
		} else {
			link.Next = l
			link = l
		}
	}
	return head
}

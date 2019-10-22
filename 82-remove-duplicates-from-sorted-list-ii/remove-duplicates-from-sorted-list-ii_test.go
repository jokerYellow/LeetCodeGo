package leetcode

import "testing"

func Test0(t *testing.T) {
	input := generateLinkList([]int{2, 2})
	output := deleteDuplicates(input)
	if checkEqualLink(output, nil) == false {
		t.Fail()
	}
}

func Test1(t *testing.T) {
	input := generateLinkList([]int{1, 2, 2, 3})
	output := deleteDuplicates(input)
	expect := generateLinkList([]int{1, 3})
	if checkEqualLink(output, expect) == false {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	input := generateLinkList([]int{1, 2, 2, 2})
	output := deleteDuplicates(input)
	expect := generateLinkList([]int{1})
	if checkEqualLink(output, expect) == false {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	input := generateLinkList([]int{1, 2, 2, 2, 2, 2, 2, 2, 3})
	output := deleteDuplicates(input)
	expect := generateLinkList([]int{1, 3})
	if checkEqualLink(output, expect) == false {
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
	if h1 == nil && h2 == nil {
		return true
	}
	return false
}

func generateLinkList(arr []int) *ListNode {
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

package leetcode

import "testing"

func Test1(t *testing.T) {
	input := generateLinkList([]int{1, 2, 3, 4})
	expect := generateLinkList([]int{4, 3, 2, 1})
	output := reverseList(input)
	if !checkEqualLink(output, expect) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	input := generateLinkList([]int{1, 2})
	expect := generateLinkList([]int{2, 1})
	output := reverseList(input)
	if !checkEqualLink(output, expect) {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	input := generateLinkList([]int{})
	expect := generateLinkList([]int{})
	output := reverseList(input)
	if !checkEqualLink(output, expect) {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	output := reverseList(nil)
	if !checkEqualLink(output, nil) {
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
	if h1 != nil || h2 != nil {
		return false
	}
	return true
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

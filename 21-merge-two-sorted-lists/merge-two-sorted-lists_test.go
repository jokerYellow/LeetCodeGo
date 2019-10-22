package leetcode

import (
	"testing"
)

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

func Test0(t *testing.T) {
	l1 := GenerateLinkList([]int{1, 2, 13})
	l2 := GenerateLinkList([]int{4, 5, 6})
	output := mergeTwoLists(l1, l2)
	expect := GenerateLinkList([]int{1, 2, 4, 5, 6, 13})
	if !checkEqualLink(output, expect) {
		t.Fail()
	}
}


func Test01(t *testing.T) {
	l1 := GenerateLinkList([]int{1, 2, 4})
	l2 := GenerateLinkList([]int{1, 3, 4})
	output := mergeTwoLists(l1, l2)
	expect := GenerateLinkList([]int{1, 1, 2, 3, 4, 4})
	if !checkEqualLink(output, expect) {
		t.Fail()
	}
}


func Test1(t *testing.T) {
	l1 := GenerateLinkList([]int{1, 2, 3})
	l2 := GenerateLinkList([]int{4, 5, 6})
	output := mergeTwoLists(l1, l2)
	expect := GenerateLinkList([]int{1, 2, 3, 4, 5, 6})
	if !checkEqualLink(output, expect) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	l2 := GenerateLinkList([]int{4, 5, 6})
	output := mergeTwoLists(nil, l2)
	expect := GenerateLinkList([]int{4, 5, 6})
	if !checkEqualLink(output, expect) {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	output := mergeTwoLists(nil, nil)
	expect := GenerateLinkList([]int{})
	if !checkEqualLink(output, expect) {
		t.Fail()
	}
}

func checkEqualLink(l1, l2 *ListNode) bool {
	h1 := l1
	h2 := l2
	if l1 == nil && l2 != nil{
		return false
	}else if l1 != nil && l2 == nil{
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
package leetcode

import "testing"

func Test1(t *testing.T) {
	head := GenerateLinkList([]int{1, 2, 2, 3, 3, 4})
	expect := GenerateLinkList([]int{1, 2, 3, 4})
	output := deleteDuplicates(head)
	if !checkEqualLink(expect, output) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	head := GenerateLinkList([]int{1,1, 2, 2, 3, 3, 4, 5, 6,6})
	expect := GenerateLinkList([]int{1, 2, 3, 4, 5, 6})
	output := deleteDuplicates(head)
	if !checkEqualLink(expect, output) {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	output := deleteDuplicates(nil)
	if !checkEqualLink(nil, output) {
		t.Fail()
	}
}


func Test4(t *testing.T) {
	head := GenerateLinkList([]int{1,1})
	expect := GenerateLinkList([]int{1})
	output := deleteDuplicates(head)
	if !checkEqualLink(expect, output) {
		t.Fail()
	}
}


func Test5(t *testing.T) {
	head := GenerateLinkList([]int{1,1,2,3,3})
	expect := GenerateLinkList([]int{1,2,3})
	output := deleteDuplicates(head)
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
	if h1 == nil && h2 == nil {
		return true
	}
	return false
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

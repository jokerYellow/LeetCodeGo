package utils

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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

func CheckEqualLink(l1, l2 *ListNode) bool {
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

func (this *ListNode) description() string {
	rt := ""
	c := this
	for c != nil {
		if rt != "" {
			rt = fmt.Sprintf("%s,%d", rt, c.Val)
		} else {
			rt = fmt.Sprintf("%d", c.Val)
		}

		c = c.Next
	}
	return rt
}

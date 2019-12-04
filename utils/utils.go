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

func EqualStrings(l, r []string) bool {
	if len(l) != len(r) {
		return false
	}
	i := 0
	for i < len(l) {
		if l[i] != r[i] {
			return false
		}
		i++
	}
	return true
}

func CheckEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if CheckEqualArr(v, b[i]) == false {
			return false
		}
	}
	return true
}

//item in a ,is alse in b
func CheckItemsEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for _, ia := range a {
		contain := false
		for _, ib := range b {
			if CheckEqualArr(ia, ib) {
				contain = true
				break
			}
		}
		if !contain {
			return false
		}
	}
	return true
}

func CheckEqualArr(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func CheckStringsEqualArr(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func Print(expect, output interface{}) {
	fmt.Printf("expect:%s\noutput:%s\n", expect, output)
}

func Contain(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

package leetcode

import "testing"

func Test1(t *testing.T) {
	head := generateLinkList([]int{1, 2, 3, 4})
	expect := 3
	output := middleNode(head).Val
	if expect != output {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	head := new(ListNode)
	expect := 0
	output := middleNode(head).Val
	if expect != output {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	output := middleNode(nil)
	if nil != output {
		t.Fail()
	}
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

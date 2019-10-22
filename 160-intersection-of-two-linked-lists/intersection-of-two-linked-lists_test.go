package leetcode

import "testing"

func Test1(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 0
	l2 := new(ListNode)
	l2.Val = 9
	l1.Next = l2
	l3 := new(ListNode)
	l3.Val = 1
	l2.Next = l3


	r1 := new(ListNode)
	r1.Val = 3

	i1 := new(ListNode)
	i1.Val = 2
	i2 := new(ListNode)
	i2.Val = 4
	i1.Next = i2

	l3.Next = i1
	r1.Next = i1

	output := getIntersectionNode(l1, r1)
	if output != i1 {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 1
	l2 := new(ListNode)
	l1.Next = l2
	l3 := new(ListNode)
	l3.Val = 12
	l2.Next = l3

	r1 := new(ListNode)
	r1.Val = 2
	r2 := new(ListNode)
	r2.Val = 3
	r1.Next = r2

	i1 := new(ListNode)
	i1.Val = 2

	l3.Next = i1
	r2.Next = i1

	output := getIntersectionNode(l1, r1)
	if output != i1 {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 1
	l2 := new(ListNode)
	l1.Next = l2
	l3 := new(ListNode)
	l3.Val = 12
	l2.Next = l3

	r1 := new(ListNode)
	r1.Val = 2
	r2 := new(ListNode)
	r2.Val = 3
	r1.Next = r2

	i1 := new(ListNode)
	i1.Val = 2

	output := getIntersectionNode(l1, r1)
	if output != nil {
		t.Fail()
	}
}
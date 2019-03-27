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
	l3.Next = l1
	if hasCycle(l1) == false {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	l1 := new(ListNode)
	if hasCycle(l1) == true {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 0
	l2 := new(ListNode)
	l2.Val = 9
	l1.Next = l2
	l3 := new(ListNode)
	l3.Val = 1
	l2.Next = l3
	if hasCycle(l1) == true {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	if hasCycle(nil) == true {
		t.Fail()
	}
}

func Test5(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 3
	l2 := new(ListNode)
	l2.Val = 2
	l1.Next = l2
	l3 := new(ListNode)
	l3.Val = 0
	l2.Next = l3
	l4 := new(ListNode)
	l4.Val = 4
	l3.Next = l4

	l4.Next = l2

	if hasCycle(l1) == false {
		t.Fail()
	}

	if hasCycle(l2) == false {
		t.Fail()
	}

	if hasCycle(l3) == false {
		t.Fail()
	}

	if hasCycle(l4) == false {
		t.Fail()
	}
}

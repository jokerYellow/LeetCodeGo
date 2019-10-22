package leetcode

import "testing"

func Test1(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 0

	l2 := new(ListNode)
	l2.Val = 9

	l3 := new(ListNode)
	l3.Val = 1

	l1.Next = l2
	l2.Next = l3
	l3.Next = l2

	output := detectCycle(l1)
	if output != l2 {
		t.Fail()
	}
}


func Test2(t *testing.T) {
	output := detectCycle(nil)
	if output != nil {
		t.Fail()
	}
}


func Test3(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 0

	l2 := new(ListNode)
	l2.Val = 9

	l3 := new(ListNode)
	l3.Val = 1

	l1.Next = l2
	l2.Next = l3

	output := detectCycle(l1)
	if output != nil {
		t.Fail()
	}
}



func Test4(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 0

	l2 := new(ListNode)
	l2.Val = 9

	l3 := new(ListNode)
	l3.Val = 1

	l1.Next = l2
	l2.Next = l3
	l3.Next = l1

	output := detectCycle(l1)
	if output != l1 {
		t.Fail()
	}
}




func Test5(t *testing.T) {
	l1 := new(ListNode)
	l1.Val = 1

	l2 := new(ListNode)
	l2.Val = 2

	l3 := new(ListNode)
	l3.Val = 3

	l4 := new(ListNode)
	l4.Val = 4
	l5 := new(ListNode)
	l5.Val = 5
	l6 := new(ListNode)
	l6.Val = 6
	l7 := new(ListNode)
	l7.Val = 7
	l8 := new(ListNode)
	l8.Val = 8
	l9 := new(ListNode)
	l9.Val = 9
	l10 := new(ListNode)
	l10.Val = 10

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	l5.Next = l6
	l6.Next = l7
	l7.Next = l8
	l8.Next = l9
	l9.Next = l10
	l10.Next = l9

	output := detectCycle(l1)
	if output != l9 {
		t.Fail()
	}
}

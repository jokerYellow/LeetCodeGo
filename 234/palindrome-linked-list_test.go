package leetcode

import "testing"

func Test0(t *testing.T) {
	head := GenerateLinkList([]int{1,2,1})
	if isPalindrome(head) == false{
		t.Fail()
	}
}


func Test1(t *testing.T) {
	head := GenerateLinkList([]int{1,2,3,4})
	if isPalindrome(head) == true{
		t.Fail()
	}
}

func Test2(t *testing.T) {
	head := GenerateLinkList([]int{1,2,2,1})
	if isPalindrome(head) == false{
		t.Fail()
	}
}

func Test3(t *testing.T) {
	head := GenerateLinkList([]int{1,2,3,3,2,1})
	if isPalindrome(head) == false{
		t.Fail()
	}
}


func Test4(t *testing.T) {
	head := GenerateLinkList([]int{1,2,3,3,3,2,1})
	if isPalindrome(head) == false{
		t.Fail()
	}
}


func Test5(t *testing.T) {
	if isPalindrome(nil) == false{
		t.Fail()
	}
}


func Test6(t *testing.T) {
	head := GenerateLinkList([]int{1})
	if isPalindrome(head) == false{
		t.Fail()
	}
}


func Test7(t *testing.T) {
	head := GenerateLinkList([]int{1,1})
	if isPalindrome(head) == false{
		t.Fail()
	}
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

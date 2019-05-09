package leetcode

import "testing"

func Test1(t *testing.T) {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	expect := 6
	if trap(input) != expect {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	var input []int
	expect := 0
	if trap(input) != expect {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	input := []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0}
	expect := 4
	if trap(input) != expect {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	input := []int{4, 2, 0, 3, 2, 5}
	expect := 9
	if trap(input) != expect {
		t.Fail()
	}
}

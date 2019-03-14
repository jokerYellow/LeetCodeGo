package leetCode

import (
	"testing"
)

func TestFunction(t *testing.T) {
	t1 := &TreeNode{2, nil, nil}

	t2 := &TreeNode{3, t1, nil}

	t3 := &TreeNode{1, nil, t2}

	output := getMinimumDifference(t3)
	expect := 1
	if output != expect {
		t.Fail()
	}
}

func TestBounds(t *testing.T) {
	t3 := &TreeNode{1, nil, nil}

	output := getMinimumDifference(t3)
	expect := 0
	if output != expect {
		t.Fail()
	}
}

func TestNegative(t *testing.T) {
	output := getMinimumDifference(nil)
	expect := 0
	if output != expect {
		t.Fail()
	}
}

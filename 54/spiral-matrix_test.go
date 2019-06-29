package _4

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

func Test1(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	output := spiralOrder(matrix)
	expect := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	if !utils.CheckEqualArr(output, expect) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	matrix := [][]int{}
	output := spiralOrder(matrix)
	expect := []int{}
	if !utils.CheckEqualArr(output, expect) {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	matrix := [][]int{{1, 2, 3, 4, 5}}
	output := spiralOrder(matrix)
	expect := []int{1, 2, 3, 4, 5}
	if !utils.CheckEqualArr(output, expect) {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	matrix := [][]int{{1}, {2}, {3}, {4}, {5}}
	output := spiralOrder(matrix)
	expect := []int{1, 2, 3, 4, 5}
	if !utils.CheckEqualArr(output, expect) {
		t.Fail()
	}
}

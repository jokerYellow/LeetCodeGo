package _9

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

func Test1(t *testing.T) {
	intput := 1
	expect := [][]int{{1}}
	if !utils.CheckEqual(expect, generateMatrix(intput)) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	intput := 0
	expect := [][]int{}
	if !utils.CheckEqual(expect, generateMatrix(intput)) {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	intput := 3
	expect := [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}
	if !utils.CheckEqual(expect, generateMatrix(intput)) {
		t.Fail()
	}
}

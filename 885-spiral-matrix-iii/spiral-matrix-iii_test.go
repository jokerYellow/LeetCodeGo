package _85_spiral_matrix_iii

import (
	"fmt"
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

func Test1(t *testing.T) {
	expect := [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}
	if !utils.CheckEqual(expect, spiralMatrixIII(1, 4, 0, 0)) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	expect := [][]int{{1, 4}, {1, 5}, {2, 5}, {2, 4}, {2, 3}, {1, 3}, {0, 3}, {0, 4}, {0, 5}, {3, 5}, {3, 4}, {3, 3}, {3, 2}, {2, 2}, {1, 2}, {0, 2}, {4, 5}, {4, 4}, {4, 3}, {4, 2}, {4, 1}, {3, 1}, {2, 1}, {1, 1}, {0, 1}, {4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 0}}
	output := spiralMatrixIII(5, 6, 1, 4)
	fmt.Printf("expect:%d\n", expect)
	fmt.Println("----")
	fmt.Printf("output:%d\n", output)
	if !utils.CheckEqual(expect, output) {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	expect := [][]int{}
	output := spiralMatrixIII(0, 0, 1, 4)
	fmt.Printf("expect:%d\n", expect)
	fmt.Println("----")
	fmt.Printf("output:%d\n", output)
	if !utils.CheckEqual(expect, output) {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	expect := [][]int{{1, 1}, {1, 0}, {0, 0}, {0, 1}}
	output := spiralMatrixIII(2, 2, 4, 4)
	fmt.Printf("expect:%d\n", expect)
	fmt.Println("----")
	fmt.Printf("output:%d\n", output)
	if !utils.CheckEqual(expect, output) {
		t.Fail()
	}
}

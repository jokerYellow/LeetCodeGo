package _5_3sum

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

func Test(t *testing.T) {
	test(t, []int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}})
	test(t, []int{-2, -1, -1, 0, 1, 2}, [][]int{{-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}})
	test(t, []int{0, 0, 0, 0}, [][]int{{0, 0, 0}})
	test(t, []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}})
	test(t, []int{}, [][]int{})
	test(t, []int{1, 2, 3}, [][]int{})
}

func test(t *testing.T, input []int, expect [][]int) {
	output := threeSum(input)
	utils.Print(expect, output)
	if !utils.CheckEqual(expect, output) {
		t.Fail()
	}
}

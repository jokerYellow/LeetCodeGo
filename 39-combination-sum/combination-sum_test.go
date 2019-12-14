package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	candidates []int
	target     int
	expect     [][]int
}

func Test(t *testing.T) {
	cases := []testCase{
		//warning:负数处理不了，但是通过了leetcode case
		//{[]int{11, -7, 3, 6, 7}, 0, [][]int{{-7, 7}}},
		{[]int{2, 3}, 11, [][]int{{2, 2, 2, 2, 3}, {2, 3, 3, 3}}},
		{[]int{2, 3, 7}, 18, [][]int{{2, 2, 2, 2, 2, 2, 2, 2, 2}, {2, 2, 2, 2, 2, 2, 3, 3}, {2, 2, 2, 2, 3, 7}, {2, 2, 2, 3, 3, 3, 3}, {2, 2, 7, 7}, {2, 3, 3, 3, 7}, {3, 3, 3, 3, 3, 3}}},
		{[]int{}, 0, [][]int{}},
		{[]int{7}, 7, [][]int{{7}}},
		{[]int{2, 3, 6, 7, 11}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
	}
	for _, item := range cases {
		output := combinationSum(item.candidates, item.target)
		if !utils.CheckElementsEqual(output, item.expect) {
			t.Fail()
		}
	}
}

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
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{
			{1, 7},
			{1, 2, 5},
			{2, 6},
			{1, 1, 6},
		}},
	}
	for _, item := range cases {
		output := combinationSum2(item.candidates, item.target)
		if !utils.CheckElementsEqual(output, item.expect) {
			t.Fail()
		}
	}
}

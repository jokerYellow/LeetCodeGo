package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	digits []int
	expect []int
}

func Test(t *testing.T) {
	cases := []testCase{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{9}, []int{1, 0}},
		{[]int{0}, []int{1}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
	}
	for _, c := range cases {
		output := plusOne(c.digits)
		if !utils.CheckEqualArr(output, c.expect) {
			utils.Print(c.expect, output)
			t.Fail()
		}
	}
}

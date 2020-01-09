package two_sum_ii_input_array_is_sorted

import (
	"fmt"
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	numbers []int
	target  int
	expect  []int
}

func Test(t *testing.T) {
	cases := []testCase{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
	}
	for _, c := range cases {
		o := twoSum(c.numbers, c.target)
		if !utils.CheckEqualArr(o, c.expect) {
			fmt.Println(c)
			utils.Print(c.expect, o)
			t.Fail()
		}
	}
}

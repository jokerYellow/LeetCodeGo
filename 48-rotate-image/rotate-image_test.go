package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	matrix [][]int
	expect [][]int
}

func Test(t *testing.T) {
	cases := []testCase{
		{[][]int{
			{1},
		},
			[][]int{
				{1},
			},
		},
		{[][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		},
			[][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
		{[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
			[][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
	}

	for _, c := range cases {
		rotate(c.matrix)
		if !utils.CheckEqual(c.matrix, c.expect) {
			utils.Print(c.expect, c.matrix)
			t.Fail()
		}
	}
}

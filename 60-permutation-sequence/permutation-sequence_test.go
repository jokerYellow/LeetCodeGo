package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	n, k   int
	expect string
}

func Test(t *testing.T) {
	cases := []testCase{
		{4, 12, "2431"},
		{3, 5, "312"},
		{3, 4, "231"},
		{3, 1, "123"},
		{3, 2, "132"},
		{2, 2, "21"},
		{1, 0, "1"},
		{1, 1, "1"},
		{3, 3, "213"},
		{4, 9, "2314"},
	}
	for _, c := range cases {
		output := getPermutation(c.n, c.k)
		if output != c.expect {
			utils.Print(c.expect, output)
			t.Fail()
		}
	}
}

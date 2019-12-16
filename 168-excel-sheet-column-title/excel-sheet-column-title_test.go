package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	n      int
	expect string
}

func Test(t *testing.T) {
	cases := []testCase{
		{52, "AZ"},
		{701, "ZY"},
		{27, "AA"},
		{1, "A"},
		{28, "AB"},
	}
	for _, c := range cases {
		output := convertToTitle(c.n)
		if output != c.expect {
			utils.Print(c.expect, output)
			t.Fail()
		}
	}
}

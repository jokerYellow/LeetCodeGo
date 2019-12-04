package leetcode

import (
	"testing"

	"github.com/jokerYellow/leetcode/utils"
)

type testCase struct {
	digits string
	expect []string
}

func Test(t *testing.T) {
	cases := []testCase{
		{"13121", []string{"da", "db", "dc", "ea", "eb", "ec", "fa", "fb", "fc"}},
		{"12131", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"2", []string{"a", "b", "c"}},
		{"", []string{}},
	}
	for _, c := range cases {
		output := letterCombinations(c.digits)
		if !utils.CheckStringsEqualArr(output, c.expect) {
			t.Fail()
			utils.Print(c.expect, output)
		}
	}

}

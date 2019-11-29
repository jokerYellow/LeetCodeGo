package leetcode

import "testing"

type testCase struct {
	s       string
	numRows int
	expect  string
}

func Test(t *testing.T) {
	cases := []testCase{
		{"", 1, ""},
		{"", 3, ""},
		{"A BC", 3, "A CB"},
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	}
	for _, c := range cases {
		output := convert(c.s, c.numRows)
		if output != c.expect {
			t.Fail()
		}
	}
}

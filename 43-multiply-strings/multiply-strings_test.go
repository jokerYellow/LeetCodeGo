package leetcode

import "testing"

type testCase struct {
	num1, num2, expect string
}

func Test(t *testing.T) {
	cases := []testCase{
		{"123", "456", "56088"},
		{"2", "3", "6"},
		{"101", "3", "303"},
		{"123", "0", "0"},
		{"123", "", ""},
		{"1111111111111111111111111111111111111111", "1", "1111111111111111111111111111111111111111"},
	}
	for _, c := range cases {
		output := multiply(c.num1, c.num2)
		if c.expect != output {
			t.Fail()
		}
	}
}

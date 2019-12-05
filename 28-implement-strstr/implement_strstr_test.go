package leetcode

import "testing"

type testCase struct {
	haystack string
	needle   string
	expect   int
}

func Test(t *testing.T) {
	cases := []testCase{
		{"", "1", -1},
		{"", "", 0},
		{"helaaalo", "laaaloo", -1},
		{"hello", "ll", 2},
		{"aaaaa", "ll", -1},
	}
	for _, item := range cases {
		output := strStr(item.haystack, item.needle)
		if output != item.expect {
			t.Fail()
		}
	}
}

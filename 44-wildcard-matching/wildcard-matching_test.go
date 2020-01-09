package _4_wildcard_matching

import (
	"fmt"
	"testing"
)

type testCase struct {
	s      string
	p      string
	expect bool
}

func Test(t *testing.T) {
	cases := []testCase{
		{"bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab",
		"b*b*ab**ba*b**b***bba",false},
		{"ho", "ho**", true},
		{"", "?", false},
		{"adceb", "*a*b", true},
		{"mississippi", "m??*ss*?i*pi", false},
		{"acdcb", "a*c?b", false},
		{"aa", "a", false},
		{"aa", "*", true},
		{"cb", "?a", false},
	}
	for _, c := range cases {
		if c.expect != isMatch(c.s, c.p) {
			fmt.Println(c)
			t.Fail()
		}
	}
}

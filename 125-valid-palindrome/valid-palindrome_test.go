package validPalindrome

import (
	"fmt"
	"testing"
)

type testCase struct {
	input  string
	expect bool
}

func Test(t *testing.T) {
	cases := []testCase{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"", true},
		{"0P", false},
	}
	for _, c := range cases {
		o := isPalindrome(c.input)
		if o != c.expect {
			fmt.Println(c.input)
			t.Fail()
		}
	}
}

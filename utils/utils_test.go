package utils

import "testing"

import "fmt"

type testCase struct {
	input  interface{}
	expect string
}

func Test(t *testing.T) {
	cases := []testCase{
		{[]byte{'0', '9'}, "09:[48 57]"},
		{[]rune{'0', '9', '黄'}, "09黄:[48 57 40644]"},
		{[]int{0, 9}, "[0 9]"},
		{[]string{"0", "9"}, "[0 9]"},
		{"91", "91"},
		{91, "91"},
	}
	for _, c := range cases {
		o := transform(c.input)
		if o != c.expect {
			fmt.Printf("expect:%s output:%s", c.expect, o)
			t.Fail()
		}
	}
}

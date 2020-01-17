package poweroftwo

import (
	"fmt"
	"testing"
)

type testCase struct {
	input  int
	expect bool
}

func Test(t *testing.T) {
	cases := []testCase{
		{3, false},
		{2, true},
		{0, false},
		{1, true},
		{4, true}}
	for _, c := range cases {
		o := isPowerOfTwo(c.input)
		if o != c.expect {
			t.Fail()
			fmt.Println(c.input)
		}
	}
}

func Test3(t *testing.T) {
	cases := []testCase{
		{1162261467, true},
		{27, true},
		{5, false},
		{9, true},
		{3, true},
		{0, false},
		{1, true}}
	for _, c := range cases {
		o := isPowerOf(c.input, 3)
		if o != c.expect {
			t.Fail()
			fmt.Println(c.input)
		}
	}
}

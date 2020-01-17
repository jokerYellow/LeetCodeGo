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

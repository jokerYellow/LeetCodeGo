package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	board  [][]byte
	word   string
	expect bool
}

func Test(t *testing.T) {
	cases := []testCase{
		{[][]byte{
			{'a'},
		},
			"a",
			true},
		{[][]byte{
		},
			"",
			false},
		{[][]byte{
		},
			"SEE",
			false},
		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		},
			"SEE",
			true},

		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		},
			"ABCCED",
			true},

		{[][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		},
			"ABCB",
			false},
	}

	for _, c := range cases {
		output := exist(c.board, c.word)
		if output != c.expect {
			utils.Print(c.expect, output)
			t.Fail()
		}
	}
}

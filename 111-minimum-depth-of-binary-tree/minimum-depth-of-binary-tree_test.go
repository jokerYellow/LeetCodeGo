package leetcode

import (
	"fmt"
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	tree   *utils.TreeNode
	expect int
}

func Test(t *testing.T) {
	cases := []testCase{
		{utils.NewTree([]interface{}{1, 2}), 2},
		{utils.NewTree([]interface{}{3, 9, 20, 1, utils.Null, 15, 7, 10}), 3},
		{utils.NewTree([]interface{}{}), 0},
		{utils.NewTree([]interface{}{3, 9, 20, utils.Null, utils.Null, 15, 7}), 2},
	}
	for _, c := range cases {
		o := minDepth(c.tree)
		if o != c.expect {
			utils.Print(c.expect,o)
			fmt.Println(utils.Description(c.tree))
			t.Fail()
		}else{
			fmt.Println(utils.Description(c.tree))
		}
	}
}

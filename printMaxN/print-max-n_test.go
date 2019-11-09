package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

func Test(t *testing.T) {
	test(t, 6, "0\n1")
}

func test(t *testing.T, input int, expect string) {
	output := printMaxN(input)
	if output != expect {
		utils.Print(expect, output)
		t.Fail()
	}
}

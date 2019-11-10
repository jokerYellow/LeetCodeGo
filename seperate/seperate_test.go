package seperate

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

func Test(t *testing.T) {
	test(t, []int{}, []int{})
	test(t, []int{1}, []int{1})
	test(t, []int{0}, []int{0})
	test(t, []int{1, 2, 3, 4}, []int{1, 3, 2, 4})
}

func test(t *testing.T, input, expect []int) {
	output := seperate(input)
	if !utils.CheckEqualArr(expect, output) {
		t.Fail()
	}
}

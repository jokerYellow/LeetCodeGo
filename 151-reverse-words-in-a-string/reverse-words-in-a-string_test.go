package leetcode

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

func Test(t *testing.T) {
	test("", "", t)
	test(" ", "", t)
	test("a", "a", t)
	test("ab", "ab", t)
	test("the sky is blue.", "blue. is sky the", t)
}

func test(input, expect string, t *testing.T) {
	output := reverseWords(input)
	utils.Print(expect, output)
	if output != expect {
		t.Fail()
	}
}

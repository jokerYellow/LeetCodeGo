package __longest_palindromic_substring

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

func Test1(t *testing.T) {
	input := "babad"
	expect := []string{"bab", "aba"}
	output := longestPalindrome(input)

	if !utils.Contain(expect, output) {
		utils.Print(expect, output)
		t.Fail()
	}
}

func Test2(t *testing.T) {
	input := "cbba"
	expect := []string{"bb"}
	output := longestPalindrome(input)

	if !utils.Contain(expect, output) {
		utils.Print(expect, output)
		t.Fail()
	}
}

func Test3(t *testing.T) {
	input := "cbbacdeffffedc"
	expect := []string{"cdeffffedc"}
	output := longestPalindrome(input)

	if !utils.Contain(expect, output) {
		utils.Print(expect, output)
		t.Fail()
	}
}

func Test4(t *testing.T) {
	input := "faaaaaaaaaaaaaaaaaaaaaaaaaaaaaaafaaaaaaf"
	expect := []string{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
	output := longestPalindrome(input)

	if !utils.Contain(expect, output) {
		utils.Print(expect, output)
		t.Fail()
	}
}

func Test5(t *testing.T) {
	input := "a"
	expect := []string{"a"}
	output := longestPalindrome(input)

	if !utils.Contain(expect, output) {
		utils.Print(expect, output)
		t.Fail()
	}
}

func Test6(t *testing.T) {
	input := ""
	expect := []string{""}
	output := longestPalindrome(input)

	if !utils.Contain(expect, output) {
		utils.Print(expect, output)
		t.Fail()
	}
}

func Test7(t *testing.T) {
	input := "ab"
	expect := []string{"b", "a"}
	output := longestPalindrome(input)

	if !utils.Contain(expect, output) {
		utils.Print(expect, output)
		t.Fail()
	}
}

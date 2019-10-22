package leetcode

import (
	"testing"
)

func TestNormal(t *testing.T) {
	input := "aba"
	output := validPalindrome(input)
	if output != true {
		t.Fail()
	}
}

func TestNormal2(t *testing.T) {
	input := "abcdefg8hijklmn"
	output := validPalindrome(input)
	if output != false {
		t.Fail()
	}
}

func TestNormal3(t *testing.T) {
	input := "aaaaaaaaaaaaaa"
	output := validPalindrome(input)
	if output != true {
		t.Fail()
	}
}

func TestNormal4(t *testing.T) {
	input := "baaaaaaakaaaaaaa"
	output := validPalindrome(input)
	if output != true {
		t.Fail()
	}
}

func TestNormal5(t *testing.T) {
	input := "abcdefgkkgfedcba"
	output := validPalindrome(input)
	if output != true {
		t.Fail()
	}
}

func TestNormal6(t *testing.T) {
	input := "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"
	output := validPalindrome(input)
	if output != true {
		t.Fail()
	}
}

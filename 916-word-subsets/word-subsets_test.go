package _16_word_subsets

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

func Test1(t *testing.T) {
	A := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	B := []string{"ec", "oc", "ceo"}
	expect := []string{"facebook", "leetcode"}
	output := wordSubsets(A, B)
	if !utils.EqualStrings(expect, output) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	A := []string{"amazonccccc", "apple", "facebook", "google", "leetcode"}
	B := []string{"a", "oc", "ceo"}
	expect := []string{"facebook"}
	output := wordSubsets(A, B)
	if !utils.EqualStrings(expect, output) {
		t.Fail()
	}
}

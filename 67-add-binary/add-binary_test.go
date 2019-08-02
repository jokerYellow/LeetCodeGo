package _7_add_binary

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

func Test1(t *testing.T) {
	a := "11"
	b := "1"
	expect := "100"
	output := addBinary(a, b)
	if expect != output {
		utils.Print(expect, output)
		t.Fail()
	}
}

func Test2(t *testing.T) {
	a := "100001"
	b := "1"
	expect := "100010"
	output := addBinary(a, b)
	if expect != output {
		utils.Print(expect, output)
		t.Fail()
	}
}

func Test3(t *testing.T) {
	a := "1"
	b := "0"
	expect := "1"
	output := addBinary(a, b)
	if expect != output {
		utils.Print(expect, output)
		t.Fail()
	}
}

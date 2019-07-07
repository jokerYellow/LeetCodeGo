package _02_remove_k_digits

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

func Test1(t *testing.T) {
	num := "1432219"
	k := 3
	expect := "1219"
	output := removeKdigits(num, k)
	if expect != output {
		utils.Print(expect, output)
		t.Fail()
	}
}

func Test2(t *testing.T) {
	num := "10200"
	k := 1
	expect := "200"
	output := removeKdigits(num, k)
	if expect != output {
		utils.Print(expect, output)
		t.Fail()
	}
}

func Test3(t *testing.T) {
	num := "10"
	k := 1
	expect := "0"
	output := removeKdigits(num, k)
	if expect != output {
		utils.Print(expect, output)
		t.Fail()
	}
}

func Test4(t *testing.T) {
	num := "149219"
	k := 3
	expect := "119"
	output := removeKdigits(num, k)
	if expect != output {
		utils.Print(expect, output)
		t.Fail()
	}
}

func Test5(t *testing.T) {
	num := "10"
	k := 2
	expect := "0"
	output := removeKdigits(num, k)
	if expect != output {
		utils.Print(expect, output)
		t.Fail()
	}
}

func Test6(t *testing.T) {
	num := "112"
	k := 2
	expect := "1"
	output := removeKdigits(num, k)
	if expect != output {
		utils.Print(expect, output)
		t.Fail()
	}
}

package leetcode

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	test("", 0, t)
	test("214748364722", 1<<31-1, t)
	test("-214748364722", -1<<31, t)
	test("    -1111", -1111, t)
	test("adwadwadw    1111", 0, t)
	test("    1111", 1111, t)
	test("1111", 1111, t)
	test("-", 0, t)
	test("+", 0, t)
	test("1", 1, t)
}

func test(input string, expect int, t *testing.T) {
	output := myAtoi(input)
	fmt.Println(input, output, expect)
	if expect != output {
		t.Fail()
	}
}

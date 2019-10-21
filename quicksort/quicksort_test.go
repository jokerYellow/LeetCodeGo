package quicksort

import (
	"fmt"
	"testing"

	"github.com/jokerYellow/leetcode/utils"
)

func TestQuicksort(t *testing.T) {
	input := []int{1, 2, 3, 5, 2}
	expect := []int{1, 2, 2, 3, 5}
	test(input, expect, t)
}

func TestQuicksort1(t *testing.T) {
	input := []int{1, 2, 3, 5, 2, 10}
	expect := []int{1, 2, 2, 3, 5, 10}
	test(input, expect, t)
	fmt.Println("hello")
}

func TestQuicksort2(t *testing.T) {
	input := []int{1}
	expect := []int{1}
	test(input, expect, t)
}

func TestQuicksort3(t *testing.T) {
	input := []int{}
	expect := []int{}
	test(input, expect, t)
}

func test(input, expect []int, t *testing.T) {
	output := quisksort(input)
	utils.Print(expect, output)
	if !utils.CheckEqualArr(expect, output) {
		t.Fail()
	}
}

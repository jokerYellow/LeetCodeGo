package _6_remove_duplicates_from_sorted_array

import (
	"fmt"
	"testing"
)

func test(input []int, expect int) bool {
	origin := input
	output := removeDuplicates(input)
	fmt.Printf("origin : %d\ninput : %d \noutput : %d\nexpect : %d\n", origin, input, output, expect)
	return expect == output
}

func Test1(t *testing.T) {
	if !test([]int{1, 2, 3}, 3) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	if !test([]int{1, 1, 1}, 1) {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	if !test([]int{1, 2, 3, 4, 4, 4, 4, 6, 7, 8}, 7) {
		t.Fail()
	}
}

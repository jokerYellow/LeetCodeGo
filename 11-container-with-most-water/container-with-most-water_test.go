package leetcode

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	test(t, []int{2, 3, 4, 5, 18, 17, 6}, 17)
	test(t, []int{2, 3, 10, 5, 7, 8, 9}, 36)
	test(t, []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49)
	test(t, []int{19, 1, 2}, 4)
	test(t, []int{1, 2}, 1)
	test(t, []int{1, 1}, 1)
}

func test(t *testing.T, height []int, expect int) {
	output := maxArea(height)
	fmt.Println(output, expect)
	if output != expect {
		t.Fail()
	}
}

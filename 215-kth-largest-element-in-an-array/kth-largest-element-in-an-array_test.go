package leetcode

import "testing"

func Test(t *testing.T) {
	test([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4, t)
	test([]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 6, 6, 6}, 4, 6, t)
	test([]int{1, 2, 3}, 1, 3, t)
	test([]int{3, 2, 1, 5, 6, 4}, 2, 5, t)
}

func test(nums []int, k int, expect int, t *testing.T) {
	if expect != findKthLargest(nums, k) {
		t.Fail()
	}
}

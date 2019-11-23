package leetcode

import (
	"sort"
)

func dominantIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}
	nums2 := make([]int, 0)
	nums2 = append(nums2, nums...)
	sort.Ints(nums2)
	if nums2[len(nums2)-1] < nums2[len(nums2)-2]*2 {
		return -1
	}
	for index, value := range nums {
		if value == nums2[len(nums2)-1] {
			return index
		}
	}
	return -1
}

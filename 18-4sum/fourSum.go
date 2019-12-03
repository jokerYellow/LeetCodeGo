package leetcode

import "sort"

/*
https://leetcode.com/problems/4sum

*/
//O(n3)
func fourSum(nums []int, target int) [][]int {
	var result [][]int
	if len(nums) < 4 {
		return result
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := len(nums) - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					l := left + 1
					for l < right && nums[l] == nums[left] {
						l++
					}
					left = l
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return result
}

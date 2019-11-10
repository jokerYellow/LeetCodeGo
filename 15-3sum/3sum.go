package _5_3sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var components [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				components = append(components, []int{nums[i], nums[l], nums[r]})
				for l < len(nums)-1 && nums[l] == nums[l+1] {
					l++
				}
				l++
			} else if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			}
		}
	}
	return components
}

package _5_3sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var components [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			opposite := - nums[i] - nums[j]
			if j > i + 1 && nums[j] == nums[j-1]{
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[k] == opposite {
					components = append(components, []int{nums[i], nums[j], nums[k]})
					break
				}
			}
		}
	}
	return components
}

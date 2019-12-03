package leetcode

import (
	"log"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		log.Fatalln("invalid input")
	}
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]
	for i, value := range nums {
		second := i + 1
		third := len(nums) - 1
		for second < third {
			sum := value + nums[second] + nums[third]
			if abs(target-sum) < abs(target-closest) {
				closest = sum
			}
			if sum < target {
				second += 1
			} else if sum == target {
				return sum
			} else {
				third -= 1
			}
		}
	}
	return closest
}

func abs(l int) int {
	if l > 0 {
		return l
	}
	return -1 * l
}

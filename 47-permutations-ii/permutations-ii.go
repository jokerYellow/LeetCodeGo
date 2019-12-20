package leetcode

import "sort"

/*
https://leetcode.com/problems/permutations-ii/
47. Permutations II
Medium

1433

50

Favorite

Share
Given a collection of numbers that might contain duplicates, return all possible unique permutations.

Example:

Input: [1,1,2]
Output:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return _permute(nums)
}

func _permute(nums []int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	for index, num := range nums {
		if index >= 1 && num == nums[index-1] {
			continue
		}
		var subNums = subArray(nums, index)
		for _, n := range _permute(subNums) {
			t := []int{num}
			t = append(t, n...)
			result = append(result, t)
		}
	}
	return result
}

func subArray(nums []int, index int) []int {
	var rt []int
	for i, num := range nums {
		if i != index {
			rt = append(rt, num)
		}
	}
	return rt
}

package leetcode

import "sort"

/*
https://leetcode.com/problems/combination-sum-ii/
40. Combination Sum II
Medium

1223

52

Favorite

Share
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

Each number in candidates may only be used once in the combination.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8,
A solution set is:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
Example 2:

Input: candidates = [2,5,2,1,2], target = 5,
A solution set is:
[
  [1,2,2],
  [5]
]40. Combination Sum II
Medium

1223

52

Favorite

Share
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

Each number in candidates may only be used once in the combination.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8,
A solution set is:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
Example 2:

Input: candidates = [2,5,2,1,2], target = 5,
A solution set is:
[
  [1,2,2],
  [5]
]
*/
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return _combinationSum2(candidates, target)
}

func _combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	if len(candidates) == 0 {
		return result
	}
	left := candidates[0]
	lefts := []int{left}
	for left < target {
		for _, item := range _combinationSum2(candidates[1:], target-left) {
			var valid []int
			valid = append(valid, lefts...)
			valid = append(valid, item...)
			result = append(result, valid)
		}
		break
	}
	if left == target {
		result = append(result, lefts)
	}
	count := 1
	for count < len(candidates) && candidates[count] == candidates[0] {
		count++
	}
	if count < len(candidates) && candidates[count] != candidates[0] {
		result = append(result, _combinationSum2(candidates[count:], target)...)
	}
	return result
}

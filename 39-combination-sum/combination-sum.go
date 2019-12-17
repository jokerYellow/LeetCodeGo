package leetcode

/*
39. Combination Sum
Medium

2735

85

Favorite

Share
Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

The same repeated number may be chosen from candidates unlimited number of times.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]
Example 2:

Input: candidates = [2,3,5], target = 8,
A solution set is:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
*/
func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	if len(candidates) == 0 {
		return result
	}
	left := candidates[0]
	lefts := []int{left}
	for left < target {
		for _, item := range combinationSum(candidates[1:], target-left) {
			var valid []int
			valid = append(valid, lefts...)
			valid = append(valid, item...)
			result = append(result, valid)
		}
		left += candidates[0]
		lefts = append(lefts, candidates[0])
	}
	if left == target {
		result = append(result, lefts)
	}
	result = append(result, combinationSum(candidates[1:], target)...)
	return result
}

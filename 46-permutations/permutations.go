package leetcode

/*
https://leetcode.com/problems/permutations/
46. Permutations
Medium

2771

89

Favorite

Share
Given a collection of distinct integers, return all possible permutations.

Example:

Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/
func permute(nums []int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	for index, num := range nums {
		var subNums = subArray(nums, index)
		for _, n := range permute(subNums) {
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

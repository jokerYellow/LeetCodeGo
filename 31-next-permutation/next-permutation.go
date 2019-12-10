package leetcode

/*
https://leetcode.com/problems/next-permutation
31. Next Permutation
Medium

2445

815

Favorite

Share
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/
func nextPermutation(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	lastIndex := len(nums) - 2
	for lastIndex >= 0 && nums[lastIndex] >= nums[lastIndex+1] {
		lastIndex--
	}
	if lastIndex >= 0 {
		minBiggerIndex := lastIndex + 1
		for minBiggerIndex < len(nums)-1 && nums[minBiggerIndex+1] > nums[lastIndex] {
			minBiggerIndex++
		}
		nums[lastIndex], nums[minBiggerIndex] = nums[minBiggerIndex], nums[lastIndex]
	}

	i := lastIndex + 1
	j := len(nums) - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}

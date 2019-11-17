package leetcode

/*
35. Search Insert Position
Easy

1627

207

Favorite

Share
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:

Input: [1,3,5,6], 5
Output: 2
Example 2:

Input: [1,3,5,6], 2
Output: 1
Example 3:

Input: [1,3,5,6], 7
Output: 4
Example 4:

Input: [1,3,5,6], 0
Output: 0
*/
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	return _searchInsert(nums, target, 0, len(nums)-1)
}

func _searchInsert(nums []int, target, start, end int) int {
	if nums[start] > target {
		return start
	} else if nums[end] < target {
		return end + 1
	}
	mid := start + (end-start)/2
	if nums[mid] < target {
		return _searchInsert(nums, target, mid+1, end)
	} else if nums[mid] > target {
		return _searchInsert(nums, target, start, mid-1)
	} else {
		return mid
	}
}

package leetcode

/*
https://leetcode.com/problems/search-in-rotated-sorted-array/
33. Search in Rotated Sorted Array
Medium

3238

374

Favorite

Share
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
*/
func search(nums []int, target int) int {
	return _search(nums, 0, len(nums)-1, target)
}

func _search(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}
	mid := start + (end-start)/2
	if nums[mid] == target {
		return mid
	}
	if nums[start] > nums[end] {
		if nums[mid] >= nums[start] {
			if nums[start] <= target && nums[mid] > target {
				return _search(nums, start, mid-1, target)
			} else {
				return _search(nums, mid+1, end, target)
			}
		} else {
			if nums[end] >= target && nums[mid] < target {
				return _search(nums, mid+1, end, target)
			} else {
				return _search(nums, start, mid-1, target)
			}
		}
	} else {
		if nums[mid] < target {
			return _search(nums, mid+1, end, target)
		} else {
			return _search(nums, start, mid-1, target)
		}
	}
}

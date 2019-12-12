package leetcode
/*
https://leetcode.com/problems/maximum-subarray/
53. Maximum Subarray
Easy

5687

239

Favorite

Share
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
 */
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	leftBound := 0
	rightBound := leftBound
	maxValue := nums[leftBound]
	var current int
	for rightBound < len(nums) {
		if leftBound == rightBound {
			current = nums[leftBound]
		} else {
			current += nums[rightBound]
		}
		if current > maxValue {
			maxValue = current
		}
		if current <= 0 {
			leftBound = rightBound + 1
			rightBound = leftBound
		} else {
			rightBound += 1
		}
	}
	return maxValue
}

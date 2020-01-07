package leetcode

/*
https://leetcode.com/problems/sliding-window-maximum/
239. Sliding Window Maximum
Hard

2414

142

Add to List

Share
Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.

Example:

Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
Output: [3,3,5,5,6,7]
Explanation:

Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
Note:
You may assume k is always valid, 1 ≤ k ≤ input array's size for non-empty array.

Follow up:
Could you solve it in linear time?
*/
func maxSlidingWindow(nums []int, k int) []int {
	var rt []int
	var t []int
	for i := 0; i < len(nums); i++ {
		for len(t) > 0 && t[0] <= i-k {
			t = t[1:]
		}

		for len(t) != 0 && nums[t[len(t)-1]] < nums[i] {
			t = t[:len(t)-1]
		}
		t = append(t, i)
		if i >= k-1 {
			rt = append(rt, nums[t[0]])
		}
	}
	return rt
}

func _maxSlidingWindow(nums []int, k int) []int {
	var rt []int
	var t []int
	for i := 0; i < len(nums); i++ {
		t = append(t, nums[i])
		if len(t) == k+1 {
			t = t[1:]
		}
		if len(t) == k {
			rt = append(rt, max(t))
		}
	}
	return rt
}

func max(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

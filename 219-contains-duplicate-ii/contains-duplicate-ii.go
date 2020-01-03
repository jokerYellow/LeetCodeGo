package leetcode
/*
https://leetcode.com/problems/contains-duplicate-ii/
219. Contains Duplicate II
Easy

644

791

Add to List

Share
Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true
Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false
 */
func containsNearbyDuplicate(nums []int, k int) bool {
	n := make(map[int]int)
	for i, v := range nums {
		if _, ok := n[v]; !ok {
			n[v] = 1
		} else {
			n[v]++
			if n[v] > 1 {
				return true
			}
		}
		if i-k >= 0 {
			lb := nums[i-k]
			n[lb]--
			if n[lb] > 0 {
				return true
			} else {
				n[lb] = 0
			}
		}
	}
	return false
}

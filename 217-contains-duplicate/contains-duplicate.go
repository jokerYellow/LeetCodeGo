package leetcode
/*
https://leetcode.com/problems/contains-duplicate/
217. Contains Duplicate
Easy

562

616

Add to List

Share
Given an array of integers, find if the array contains any duplicates.

Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.

Example 1:

Input: [1,2,3,1]
Output: true
Example 2:

Input: [1,2,3,4]
Output: false
Example 3:

Input: [1,1,1,3,3,4,3,2,4,2]
Output: true
 */
func containsDuplicate(nums []int) bool {
	n := map[int]bool{}
	for _, i := range nums {
		if _, ok := n[i]; ok {
			return true
		}
		n[i] = true
	}
	return false
}

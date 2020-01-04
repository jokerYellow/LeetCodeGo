package leetcode

/*
https://leetcode.com/problems/plus-one/
66. Plus One
Easy

1137

1927

Add to List

Share
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
*/
//did not change input point,initial p = 1 is perfect
func plusOne(digits []int) []int {
	length := len(digits)
	result := make([]int, length+1)
	p := 1
	for i := length - 1; i >= 0; i-- {
		t := digits[i] + p
		p = t / 10
		result[i+1] = t % 10
	}
	if p > 0 {
		result[0] = p
	} else {
		result = result[1:]
	}
	return result
}

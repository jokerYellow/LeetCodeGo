package leetcode

/*
https://leetcode.com/problems/valid-perfect-square
367. Valid Perfect Square
Easy

576

131

Add to List

Share
Given a positive integer num, write a function which returns True if num is a perfect square else False.

Note: Do not use any built-in library function such as sqrt.

Example 1:

Input: 16
Output: true
Example 2:

Input: 14
Output: false
*/
func isPerfectSquare(num int) bool {
	top := num
	bottom := 0
	var result bool
	for {
		i := (top + bottom + 1) / 2
		t := i * i
		t1 := (i + 1) * (i + 1)
		if t <= num && t1 > num {
			result = t == num
			break
		} else if t < num {
			bottom = i
		} else if t > num {
			top = i
		}
	}
	return result
}

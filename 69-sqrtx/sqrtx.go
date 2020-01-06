package leetcode

/*
https://leetcode.com/problems/sqrtx/
69. Sqrt(x)
Easy

996

1632

Add to List

Share
Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

Input: 4
Output: 2
Example 2:

Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since
             the decimal part is truncated, 2 is returned.
*/
func mySqrt(x int) int {
	top := x
	bottom := 0
	result := 0
	for {
		i := (top + bottom + 1) / 2
		t := i * i
		t1 := (i + 1) * (i + 1)
		if t <= x && t1 > x {
			result = i
			break
		} else if t < x {
			bottom = i
		} else if t > x {
			top = i
		}
	}
	return result
}

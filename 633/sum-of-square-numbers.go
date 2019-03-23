package leetcode

import "math"

/*
633. Sum of Square Numbers
Easy

306

210

Favorite

Share
Given a non-negative integer c, your task is to decide whether there're two integers a and b such that a2 + b2 = c.

Example 1:

Input: 5
Output: True
Explanation: 1 * 1 + 2 * 2 = 5


Example 2:

Input: 3
Output: False
*/

func judgeSquareSum(c int) bool {
	i := 0
	for i <= c {
		s := i * i
		if s > c {
			return false
		}
		if canSquare(c - s) {
			return true
		}
		i++
	}
	return false
}

func canSquare(a int) bool {
	t := int(math.Sqrt(float64(a)))
	return t*t == a
}

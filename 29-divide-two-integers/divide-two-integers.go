package leetcode

/*
29. Divide Two Integers
Medium

862

4144

Favorite

Share
Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero.

Example 1:

Input: dividend = 10, divisor = 3
Output: 3
Example 2:

Input: dividend = 7, divisor = -3
Output: -2
Note:

Both dividend and divisor will be 32-bit signed integers.
The divisor will never be 0.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.
*/
func divide(dividend int, divisor int) int {
	result := _divide(dividend, divisor)
	max := 1<<31 - 1
	if result > max {
		result = max
	}
	return result
}
func _divide(dividend int, divisor int) int {
	dividendSign := 1
	if dividend < 0 {
		dividendSign = -1
		dividend = - dividend
	}
	divisorSign := 1
	if divisor < 0 {
		divisorSign = -1
		divisor = - divisor
	}
	sign := 1
	if (divisorSign ^ dividendSign) != 0 {
		sign = -1
	}
	if dividend < divisor {
		return 0
	}
	quotient := 1
	last := 0
	value := divisor
	if dividend > divisor {
		for {
			if value<<1 < dividend {
				value = value << 1
				quotient = quotient << 1
			} else {
				last = dividend - value
				break
			}
		}
	}
	return sign * (quotient + divide(last, divisor))
}

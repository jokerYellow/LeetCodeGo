package powx_n
/*
https://leetcode.com/problems/powx-n/
50. Pow(x, n)
Medium

1078

2534

Add to List

Share
Implement pow(x, n), which calculates x raised to the power n (xn).

Example 1:

Input: 2.00000, 10
Output: 1024.00000
Example 2:

Input: 2.10000, 3
Output: 9.26100
Example 3:

Input: 2.00000, -2
Output: 0.25000
Explanation: 2-2 = 1/22 = 1/4 = 0.25
Note:

-100.0 < x < 100.0
n is a 32-bit signed integer, within the range [−231, 231 − 1]
 */
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if x == 0 || x == 1 || n == 1 {
		return x
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	var result float64 = x
	i := 1
	for 2*i <= n {
		result = result * result
		i = 2 * i
	}
	return result * myPow(x, n-i)
}

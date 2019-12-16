package leetcode

/*
https://leetcode.com/problems/climbing-stairs/
70. Climbing Stairs
Easy

2988

101

Favorite

Share
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.

Example 1:

Input: 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
*/
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	one := 1
	two := 2
	for i := 3; i < n; i++ {
		t := one + two
		one = two
		two = t
	}
	return one + two
}

//Recursive,when n == 45,leetcode return time limit exceeded
func _climbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	return _climbStairs(n-1) + _climbStairs(n-2)
}

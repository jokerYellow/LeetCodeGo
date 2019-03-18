package leetcode

/*
https://leetcode.com/problems/super-pow/

Medium

115

515

Favorite

Share
Your task is to calculate ab mod 1337 where a is a positive integer and b is an extremely large positive integer given in the form of an array.

Example 1:

Input: a = 2, b = [3]
Output: 8
Example 2:

Input: a = 2, b = [1,0]
Output: 1024
(ab) mod m = ((a mod m)* (b mod m)) mod m
*/

func superPow(a int, b []int) int {
	mod := 1337
	rt := 1
	for i, v := range b {
		if i == 0 {
			rt = (rt * powerAndMod(a, v, mod)) % mod
		} else {
			rt = ((powerAndMod(rt, 10, mod) % mod) * (powerAndMod(a, v, mod) % mod)) % mod
		}
	}
	return rt
}

func powerAndMod(a, p, mod int) int {
	r := 1
	for p > 0 {
		r = (r * a) % mod
		p--
	}
	return r
}

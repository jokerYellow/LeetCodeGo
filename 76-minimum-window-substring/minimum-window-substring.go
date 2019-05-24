package _6_minimum_window_substring

/*
https://leetcode.com/problems/minimum-window-substring/
76. Minimum Window Substring
Hard

2226

156

Favorite

Share
Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

Example:

Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
Note:

If there is no such window in S that covers all characters in T, return the empty string "".
If there is such window, you are guaranteed that there will always be only one unique minimum window in S.
*/

func minWindow(s string, t string) string {
	var left, right int

	contain := func(substring, window string) bool {
		origins := map[rune]int{}
		for _, v := range substring {
			origins[v] += 1
		}
		for _, v := range window {
			if _, ok := origins[v]; ok {
				origins[v] -= 1
			}
		}
		for _, v := range origins {
			if v > 0 {
				return false
			}
		}
		return true
	}
	minSub := ""
	for right <= len(s) {
		for contain(t, s[left:right]) {
			if minSub == "" || len(minSub) > right-left {
				minSub = s[left:right]
			}
			left++
		}
		right++
	}

	return minSub
}

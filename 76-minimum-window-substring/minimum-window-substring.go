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
	return minWindow3(s, t)
}

//too slow
func minWindow1(s string, t string) string {
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

func minWindow2(s string, t string) string {
	tCount := map[rune]int{}
	for _, v := range t {
		tCount[v] += 1
	}

	isOK := func() bool {
		isOK := true
		for _, v := range tCount {
			if v > 0 {
				isOK = false
				break
			}
		}
		return isOK
	}

	var left, right int
	var min string
	for right < len(s) {
		tmp := rune(s[right])
		if _, ok := tCount[tmp]; ok {
			tCount[tmp]--
		}
		if isOK() == false {
			right++
			continue
		}
		for isOK() {
			if min == "" || len(min) > len(s[left:right+1]) {
				min = s[left : right+1]
			}
			r := rune(s[left])
			if _, ok := tCount[r]; ok {
				tCount[r]++
			}
			left++
		}
		right++
	}
	return min
}

func minWindow3(s string, t string) string {
	var left int
	min := len(s) + 1
	out := ""
	tCount := map[rune]int{}
	for _, v := range t {
		tCount[v] ++
	}
	count := 0
	for i, v := range s {
		if _, ok := tCount[v]; ok {
			tCount[v]--
			if tCount[v] >= 0 {
				count++
			}
		}
		for count == len(t) {
			if min > i-left {
				min = i - left
				out = s[left : left+min+1]
			}
			r := rune(s[left])
			if _, ok := tCount[r]; ok {
				tCount[r]++
				if tCount[r] > 0 {
					count--
				}
			}
			left++
		}
	}
	return out
}

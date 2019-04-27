package leetcode

/*
https://leetcode.com/problems/longest-substring-with-at-least-k-repeating-characters/

395. Longest Substring with At Least K Repeating Characters
Medium

688

57

Favorite

Share
Find the length of the longest substring T of a given string (consists of lowercase letters only) such that every character in T appears no less than k times.

Example 1:

Input:
s = "aaabb", k = 3

Output:
3

The longest substring is "aaa", as 'a' is repeated 3 times.
Example 2:

Input:
s = "ababbc", k = 2

Output:
5

The longest substring is "ababb", as 'a' is repeated 2 times and 'b' is repeated 3 times.
 */

func longestSubstring(s string, k int) int {
	if k == 1 {
		return len(s)
	}
	counts := map[rune]int{}
	for _, v := range s {
		counts[v] += 1
	}

	isOK := true
	for _, v := range counts {
		if v < k {
			isOK = false
		}
	}
	if isOK == true {
		return len(s)
	}
	
	var t []rune
	var maxCount int
	for _, v := range s {
		if counts[v] < k {
			count := longestSubstring(string(t), k)
			if maxCount < count {
				maxCount = count
			}
			t = []rune{}
		} else {
			t = append(t, v)
		}
	}
	tLength := longestSubstring(string(t), k)
	if tLength > maxCount {
		maxCount = tLength
	}

	return maxCount
}

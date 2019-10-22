package leetcode

/*
https://leetcode.com/problems/longest-substring-without-repeating-characters/
3. Longest Substring Without Repeating Characters
Medium

5340

282

Favorite

Share
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

"pwwkew"

p 1
w 3
k 1
e 1

edwa awf w

 */

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	info := map[rune]int{}
	max := 1
	start := 0
	for i, v := range s {
		if _, ok := info[v]; ok {
			length := i - start
			if length > max {
				max = length
			}
			for start <= info[v] {
				info[rune(s[start])] = 0
				start += 1
			}
		}
		info[v] = i
	}

	if len(s)-start > max {
		max = len(s) - start
	}
	return max
}

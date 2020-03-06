package _4_wildcard_matching

/*
https://leetcode.com/problems/wildcard-matching/
44. Wildcard Matching
Hard

1490

85

Add to List

Share
Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*'.

'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like ? or *.
Example 1:

Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input:
s = "aa"
p = "*"
Output: true
Explanation: '*' matches any sequence.
Example 3:

Input:
s = "cb"
p = "?a"
Output: false
Explanation: '?' matches 'c', but the second letter is 'a', which does not match 'b'.
Example 4:

Input:
s = "adceb"
p = "*a*b"
Output: true
Explanation: The first '*' matches the empty sequence, while the second '*' matches the substring "dce".
Example 5:

Input:
s = "acdcb"
p = "a*c?b"
Output: false
*/
//http://yucoding.blogspot.com/2013/02/leetcode-question-123-wildcard-matching.html
func isMatch(s string, p string) bool {
	i, j := 0, 0
	starPosition := -1
	sPosition := 0
	for i < len(s) {
		sc := s[i]
		var pc uint8
		if j < len(p) {
			pc = p[j]
		}
		if sc == pc || pc == '?' {
			i++
			j++
		} else if pc == '*' {
			starPosition = j
			sPosition = i
			j++
		} else if starPosition >= 0 {
			j = starPosition + 1
			i = sPosition + 1
			sPosition++
		} else {
			return false
		}
	}
	for j < len(p) && p[j] == '*' {
		j++
	}
	return j == len(p)
}

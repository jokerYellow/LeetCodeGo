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
func isMatch(s string, p string) bool {
	//fmt.Println(s, p)
	if p == "*" {
		return true
	}
	if s == p {
		return true
	}
	if p == "" {
		return false
	}
	if p == "?" && len(s) == 1 {
		return true
	}
	if p[0] == '?' {
		if len(s) == 0 {
			return false
		}
		return isMatch(s[1:], p[1:])
	} else if p[0] == '*' {
		for i := 0; i < len(s); i++ {
			if isMatch(s[i:], p[1:]) {
				return true
			}
		}
		if len(s) == 0 {
			return isMatch(s, p[1:])
		}
		return false
	} else if isPureString(p) {
		return s == p
	} else {
		for i := 0; i < len(p); i++ {
			if len(s) < i {
				return false
			}
			if p[i] == '*' || p[i] == '?' {
				return isMatch(s[i:], p[i:])
			}
			if len(s) > i && s[i] != p[i] {
				return false
			}
		}
	}
	return false
}

func isPureString(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c == '?' || c == '*' {
			return false
		}
	}
	return true
}

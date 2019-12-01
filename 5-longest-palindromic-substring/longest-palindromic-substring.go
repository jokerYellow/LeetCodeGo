package __longest_palindromic_substring

/*
https://leetcode.com/problems/longest-palindromic-substring/
5. Longest Palindromic Substring
Medium

3863

370

Favorite

Share
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
*/

func longestPalindrome(s string) string {
	normal := []byte(s)
	reversed := reverse(normal)
	var longest []byte
	length := len(s)
	for i, _ := range reversed {
		for j, _ := range normal {
			var current []byte
			if reversed[i] == normal[j] {
				for k := 0; j+k < length && i+k < length; k++ {
					if reversed[i+k] == normal[j+k] {
						current = append(current, reversed[i+k])
						if len(current) > len(longest) && isPalindrome(current) {
							longest = current
						}
					} else {
						break
					}
				}
			}
		}
	}
	return string(longest)
}

func isPalindrome(s []byte) bool {
	i := 0
	j := len(s) - 1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func reverse(bytes []byte) []byte {
	rt := []byte{}
	for j := len(bytes) - 1; j >= 0; j-- {
		rt = append(rt, bytes[j])
	}
	return rt
}

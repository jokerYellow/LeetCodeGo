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
	return longestPalindrome_expandCenter(s)
}

//1,2,1
//index:1
//length:3
//1-3/2,1+3/2
//1,2,2,1
//index:1,length:4
//1-4/2+1,1+4/2
func longestPalindrome_expandCenter(s string) string {
	if len(s) < 2 {
		return s
	}
	var longest, l, r int
	for i := 0; i < len(s); i++ {
		length := max(expandCenter(s, i, i), expandCenter(s, i, i+1))
		if length > longest {
			left := i - (length-1)/2
			right := i + length/2
			l = left
			r = right
			longest = length
		}
	}
	return s[l : r+1]
}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

func expandCenter(s string, left, right int) int {
	length := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		length = right - left + 1
		left--
		right++
	}
	return length
}

//dynamic programming
func longestPalindrome_dp(s string) string {
	length := len(s)
	var res string
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}
	for i := 0; i < length; i++ {
		for j := i; j >= 0; j-- {
			if s[i] == s[j] && i-j < 3 {
				dp[i][j] = true
			} else if s[i] == s[j] && dp[i-1][j+1] {
				dp[i][j] = true
			}
			if dp[i][j] && i-j+1 > len(res) {
				res = s[j : i+1]
			}
		}
	}
	return res
}

func _longestPalindrome(s string) string {
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

package leetcode

/*
https://leetcode.com/problems/valid-palindrome-ii/
680. Valid Palindrome II
Easy

643

42

Favorite

Share
Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.

A palindrome is a word or a phrase that is the same whether you read it backward or forward, for example, the word "refer."

Example 1:

Input: "aba"
Output: True
Example 2:

Input: "abca"
Output: True
Explanation: You could delete the character 'c'.
Note:

The string will only contain lowercase characters a-z. The maximum length of the string is 50000.
*/

func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for {
		if i >= j {
			break
		}
		left := s[i]
		right := s[j]
		if left != right {
			rt := false
			if s[i+1] == right {
				rt = rt || strictValidPalindrome(s, i+1, j)
			}
			if s[j-1] == left {
				rt = rt || strictValidPalindrome(s, i, j-1)
			}
			return rt
		}
		i++
		j--
	}
	return true
}

func strictValidPalindrome(s string, leftBound, rightBound int) bool {
	for {
		if leftBound >= rightBound {
			break
		}
		if s[leftBound] != s[rightBound] {
			return false
		}
		leftBound++
		rightBound--
	}
	return true
}

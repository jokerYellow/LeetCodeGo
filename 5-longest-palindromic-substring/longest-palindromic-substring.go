package __longest_palindromic_substring

import (
	"fmt"
)

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

type record struct {
	s string
	b bool
}

func longestPalindrome(s string) string {
	info := map[string]record{}
	l := checkAndlongestPalindrome(s, &info)
	return l
}

func checkAndlongestPalindrome(s string, info *map[string]record) (rt string) {
	fmt.Println(s)
	if palindrome(s) {
		(*info)[s] = record{s, true}
		return s
	}
	l1 := s[0 : len(s)-1]
	l2 := s[1:]
	t1 := palindrome(l1)
	t2 := palindrome(l2)
	if t1 {
		(*info)[l1] = record{l1, true}
		return l1
	} else if t2 {
		(*info)[l2] = record{l2, true}
		return l2
	} else {
		p1 := checkAndlongestPalindrome(l1, info)
		p2 := checkAndlongestPalindrome(l2, info)
		if len(p1) > len(p2) {
			(*info)[l2] = record{l2, true}
			return p1
		} else {

			return p2
		}
	}
}

func palindrome(s string) bool {
	l := 0
	r := len(s) - 1
	for {
		if l >= r {
			return true
		}
		if s[l] == s[r] && l < r {
			l++
			r--
		} else {
			return false
		}
	}
}

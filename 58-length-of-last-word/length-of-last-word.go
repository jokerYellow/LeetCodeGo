package leetcode

/*
https://leetcode.com/problems/length-of-last-word/
58. Length of Last Word
Easy

478

1977

Favorite

Share
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

Example:

Input: "Hello World"
Output: 5
*/
func lengthOfLastWord(s string) int {
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && count > 0 {
			break
		} else if s[i] != ' ' {
			count++
		}
	}
	return count
}

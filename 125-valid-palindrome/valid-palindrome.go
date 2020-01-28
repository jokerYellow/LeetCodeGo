package validPalindrome

/*
125. Valid Palindrome
Easy

865

2412

Add to List

Share
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Example 2:

Input: "race a car"
Output: false
*/

func isPalindrome(s string) bool {
	ss := []rune{}
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			ss = append(ss, c)
		} else if c >= 'A' && c <= 'Z' {
			ss = append(ss, c-'A'+'a')
		} else if c >= '0' && c <= '9' {
			ss = append(ss, c)
		}
	}
	return _isPalindrome(ss)
}

func _isPalindrome(ss []rune) bool {
	i := 0
	j := len(ss) - 1
	for i <= j {
		if ss[i] != ss[j] {
			return false
		}
		i++
		j--
	}
	return true
}

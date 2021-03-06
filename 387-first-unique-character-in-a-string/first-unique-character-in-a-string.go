package leetcode

/*
https://leetcode.com/problems/first-unique-character-in-a-string/

387. First Unique Character in a String
Easy

933

75

Favorite

Share
Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.

Examples:

s = "leetcode"
return 0.

s = "loveleetcode",
return 2.
Note: You may assume the string contain only lowercase letters.
 */

func firstUniqChar(s string) int {
	info := map[rune]int{}
	for _,v := range s{
		info[v] += 1
	}
	for i,v := range s{
		if info[v] == 1{
			return i
		}
	}
	return -1
}

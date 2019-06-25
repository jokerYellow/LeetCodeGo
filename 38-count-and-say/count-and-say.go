package _8_count_and_say

import (
	"fmt"
	"strconv"
)

/*
https://leetcode.com/problems/count-and-say/
38. Count and Say
Easy

797

6126

Favorite

Share
The count-and-say sequence is the sequence of integers with the first five terms as following:

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.

Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence.

Note: Each term of the sequence of integers will be represented as a string.



Example 1:

Input: 1
Output: "1"
Example 2:

Input: 4
Output: "1211"
*/
func countAndSay(n int) string {
	origin := ""
	for i := 1; i <= n; i++ {
		if i == 1 {
			origin = "1"
		} else {
			origin = say(origin)
		}
		fmt.Println(origin)
	}
	return origin
}

func say(s string) string {
	rt := ""
	if len(s) == 0 {
		return ""
	}
	current := rune(s[0])
	count := 1
	for _, c := range s[1:] {
		if current != c {
			rt = rt + strconv.Itoa(count) + string(current)
			current = c
			count = 1
		} else {
			count++
		}
	}
	if count > 0 {
		rt = rt + strconv.Itoa(count) + string(current)
	}
	return rt
}

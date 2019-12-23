package leetcode

import "bytes"
/*
https://leetcode.com/problems/to-lower-case/
709. To Lower Case
Easy

368

1267

Add to List

Share
Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.



Example 1:

Input: "Hello"
Output: "hello"
Example 2:

Input: "here"
Output: "here"
Example 3:

Input: "LOVELY"
Output: "lovely"
 */
func toLowerCase(str string) string {
	rt := bytes.Buffer{}
	offset := 'A' - 'a'
	for _, item := range str {
		if item >= 'A' && item <= 'Z' {
			rt.WriteRune(item - offset)
		} else {
			rt.WriteRune(item)
		}
	}
	return rt.String()
}

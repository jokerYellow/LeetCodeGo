package leetcode

/*
https://leetcode.com/problems/excel-sheet-column-title/
168. Excel Sheet Column Title
Easy

896

186

Favorite

Share
Given a positive integer, return its corresponding column title as appear in an Excel sheet.

For example:

    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB
    ...
Example 1:

Input: 1
Output: "A"
Example 2:

Input: 28
Output: "AB"
Example 3:

Input: 701
Output: "ZY"
*/
func convertToTitle(n int) string {
	return string(_convertToTitle(n))
}

func _convertToTitle(n int) []byte {
	if n == 0 {
		return []byte{}
	}
	value := n % 26
	off := n / 26
	c := byte('A' + value - 1)
	if value == 0 {
		c = 'Z'
		off--
	}
	bytes := _convertToTitle(off)
	bytes = append(bytes, c)
	return bytes
}

package _7_add_binary

/*
https://leetcode.com/problems/add-binary/

67. Add Binary
Easy

996

198

Favorite

Share
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
*/

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	cv := 0
	var bytes []byte

	for i >= 0 || j >= 0 || cv > 0 {
		vi := 0
		if i >= 0 {
			if a[i] == '1' {
				vi = 1
			} else {
				vi = 0
			}
		}
		vj := 0
		if j >= 0 {
			if b[j] == '1' {
				vj = 1
			} else {
				vj = 0
			}
		}
		sum := vi + vj + cv
		if sum%2 == 1 {
			bytes = append(bytes, '1')
		} else {
			bytes = append(bytes, '0')
		}
		cv = sum / 2
		i--
		j--
	}
	return string(reverseBytes(bytes))
}

func reverseBytes(origin []byte) []byte {
	length := len(origin)
	rt := make([]byte, len(origin))
	for i, v := range origin {
		rt[length-i-1] = v
	}
	return rt
}

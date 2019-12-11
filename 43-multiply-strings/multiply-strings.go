package leetcode

/*
https://leetcode.com/problems/multiply-strings/
43. Multiply Strings
Medium

1322

619

Favorite

Share
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Example 1:

Input: num1 = "2", num2 = "3"
Output: "6"
Example 2:

Input: num1 = "123", num2 = "456"
Output: "56088"
Note:

The length of both num1 and num2 is < 110.
Both num1 and num2 contain only digits 0-9.
Both num1 and num2 do not contain any leading zero, except the number 0 itself.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/
import "bytes"

func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return ""
	}
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	length := len(num1) + len(num2)
	p := make([]rune, length)
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			v := (num1[i] - '0') * (num2[j] - '0')
			p[i+j+1] += int32(v % 10)
			p[i+j] += int32(v / 10)
		}
	}
	for i := length - 1; i >= 1; i-- {
		p[i-1] += p[i] / 10
		p[i] = p[i] % 10
	}
	buffer := bytes.Buffer{}
	valid := false
	for _, item := range p {
		if item != 0 || valid {
			valid = true
			buffer.WriteRune(item + '0')
		}
	}
	return buffer.String()
}

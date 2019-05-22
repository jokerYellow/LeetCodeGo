package _2_integer_to_roman

import "fmt"

/*
https://leetcode.com/problems/integer-to-roman/
12. Integer to Roman
Medium

568

1826

Favorite

Share
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.

Example 1:

Input: 3
Output: "III"
Example 2:

Input: 4
Output: "IV"
Example 3:

Input: 9
Output: "IX"
Example 4:

Input: 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.
Example 5:

Input: 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/

func gewei(num int) string {
	switch num {
	case 0:
		return ""
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6:
		return "VI"
	case 7:
		return "VII"
	case 8:
		return "VIII"
	case 9:
		return "IX"
	}
	return ""
}

func shiwei(num int) string {
	switch num {
	case 0:
		return ""
	case 1:
		return "X"
	case 2:
		return "XX"
	case 3:
		return "XXX"
	case 4:
		return "XL"
	case 5:
		return "L"
	case 6:
		return "LX"
	case 7:
		return "LXX"
	case 8:
		return "LXXX"
	case 9:
		return "XC"
	}
	return ""
}

func baiwei(num int) string {
	switch num {
	case 0:
		return ""
	case 1:
		return "C"
	case 2:
		return "CC"
	case 3:
		return "CCC"
	case 4:
		return "CD"
	case 5:
		return "D"
	case 6:
		return "DC"
	case 7:
		return "DCC"
	case 8:
		return "DCCC"
	case 9:
		return "CM"
	}
	return ""
}

func qianwei(num int) string {
	switch num {
	case 0:
		return ""
	case 1:
		return "M"
	case 2:
		return "MM"
	case 3:
		return "MMM"
	}
	return ""
}

func intToRoman(num int) string {
	geweiValue := num % 10
	num = num - geweiValue
	shiweiValue := num % 100
	num = num - shiweiValue
	baiweiValue := num % 1000
	num = num - baiweiValue
	qianweiValue := num

	geweiString := gewei(geweiValue)
	shiweiString := shiwei(shiweiValue / 10)
	baiweiString := baiwei(baiweiValue / 100)
	qianweiString := qianwei(qianweiValue / 1000)
	return fmt.Sprintf("%s%s%s%s", qianweiString, baiweiString, shiweiString, geweiString)
}

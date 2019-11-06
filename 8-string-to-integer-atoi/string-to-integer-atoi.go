package leetcode

import "bytes"

func myAtoi(str string) int {
	var buffer bytes.Buffer

	for _, item := range str {
		if item == ' ' && buffer.Len() == 0 {
			continue
		}
		buffer.WriteRune(item)
	}
	if str == "+" || str == "-" {
		return 0
	}
	coreStr := buffer.String()
	max := 1<<31 - 1
	min := -1 << 31
	num := 0
	minus := 1
	for index, item := range coreStr {
		if index == 0 {
			if item == '-' {
				minus = -1
				continue
			} else if item == '+' {
				minus = 1
				continue
			}
		}
		if !isValid(item) {
			break
		}
		v := intValue(item)
		num = 10*num + v
		if minus == -1 && -1*num <= min {
			return min
		} else if minus == 1 && num >= max {
			return max
		}
	}
	return num * minus
}

func isValid(c rune) bool {
	return c >= '0' && c <= '9'
}

func intValue(c rune) int {
	return int(c) - '0'
}

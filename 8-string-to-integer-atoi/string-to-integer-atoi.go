package leetcode

func myAtoi(str string) int {
	max := (1 << 31) - 1
	min := -(1 << 31)
	num := 0
	sign := 1
	space := 0
	for index, item := range str {
		if item == ' ' {
			if index == space {
				space++
				continue
			} else {
				break
			}
		}
		if index == space {
			if item == '-' {
				sign = -1
				continue
			} else if item == '+' {
				sign = 1
				continue
			}
		}
		if !isValid(item) {
			break
		}
		v := intValue(item)
		num = 10*num + v
		if sign == -1 && -1*num <= min {
			return min
		} else if sign == 1 && num >= max {
			return max
		}
	}
	return num * sign
}

func isValid(c rune) bool {
	return c >= '0' && c <= '9'
}

func intValue(c rune) int {
	return int(c) - '0'
}

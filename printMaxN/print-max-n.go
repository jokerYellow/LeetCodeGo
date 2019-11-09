package leetcode

import (
	"bytes"
	"fmt"
)

func printMaxN(n int) string {
	count := make([]int, n)
	buffer := bytes.Buffer{}
	for increase(count) {
		buffer.WriteString(transToString(count))
		buffer.WriteByte('\n')
	}
	return buffer.String()
}

func transToString(num []int) string {
	rt := bytes.Buffer{}
	valid := false
	for _, v := range num {
		if v != 0 && !valid {
			valid = true
		}
		if valid {
			rt.WriteString(fmt.Sprintf("%d", v))
		}
	}
	return rt.String()
}

func increase(num []int) bool {
	more := 1
	for index := len(num) - 1; index >= 0; index-- {
		value := num[index]
		if (value + more) <= 9 {
			num[index] = value + more
			more = 0
			break
		} else {
			num[index] = 0
			more = 1
		}
	}
	return more != 1
}

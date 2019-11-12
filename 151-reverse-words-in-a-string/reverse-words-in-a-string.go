package leetcode

import (
	"bytes"
)

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := [][]rune{}
	var t []rune
	for _, r := range s {
		if r == ' ' && len(t) > 0 {
			runes = append(runes, t)
			t = []rune{}
		} else if r != ' ' {
			t = append(t, r)
		}
	}
	if len(t) > 0 {
		runes = append(runes, t)
	}
	buffer := bytes.Buffer{}
	for i := len(runes) - 1; i >= 0; i-- {
		buffer.WriteString(string(runes[i]))
		if i > 0 {
			buffer.WriteByte(' ')
		}
	}
	return buffer.String()
}

package leetcode

import "bytes"

func convert(s string, numRows int) string {
	if numRows == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	length := len(s)
	numLine := (length / (numRows*2 - 2)) * (numRows - 1)
	more := length % (numRows*2 - 2)
	if more <= numRows && more > 0 {
		numLine += 1
	} else if more > numRows {
		numLine += more - numRows + 1
	}
	arr := make([][]rune, numRows)
	for i := 0; i < numRows; i++ {
		arr[i] = make([]rune, numLine)
	}
	for i, c := range s {
		offset := i / (numRows*2 - 2)
		more := i % (numRows*2 - 2)
		row := 0
		number := offset * (numRows - 1)
		if more < numRows {
			row = more
		} else {
			row = numRows - (more+1)%numRows - 1
			number += (more + 1) % numRows
		}
		arr[row][number] = c
	}
	buffer := bytes.Buffer{}
	for _, item := range arr {
		for _, c := range item {
			if c != 0 {
				buffer.WriteRune(c)
			}
		}
	}
	return buffer.String()
}

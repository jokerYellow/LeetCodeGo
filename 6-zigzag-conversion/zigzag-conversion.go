package leetcode

import "bytes"

func convert(s string, numRows int) string {
	if numRows == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	arr := make([]bytes.Buffer, numRows)
	for i := 0; i < numRows; i++ {
		arr[i] = bytes.Buffer{}
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
		arr[row].WriteRune(c)
	}
	buffer := bytes.Buffer{}
	for _, item := range arr {
		buffer.Write(item.Bytes())
	}
	return buffer.String()
}

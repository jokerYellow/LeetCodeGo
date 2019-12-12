package leetcode

const size int = 9

func isValidSudoku(board [][]byte) bool {

	if len(board) != size {
		return false
	}
	for i := 0; i < size; i++ {
		if !isValidNine(board[i]) {
			return false
		}
		columnNums := make([]byte, size)
		for j := 0; j < size; j++ {
			columnNums[j] = board[j][i]
		}
		if !isValidNine(columnNums) {
			return false
		}
		rowLeftTop := (i / 3) * 3
		columnLeftTop := (i % 3) * 3
		subBox := make([]byte, size)
		for j := 0; j < size; j++ {
			row := rowLeftTop + (j / 3)
			column := columnLeftTop + (j % 3)
			subBox[j] = board[row][column]
		}
		if !isValidNine(subBox) {
			return false
		}
	}
	return true
}

func isValidNine(nums []byte) bool {
	numberMap := map[byte]bool{}
	for _, b := range nums {
		if b >= '1' && b <= '9' {
			if numberMap[b] {
				return false
			}
			numberMap[b] = true
		}
	}
	return true
}

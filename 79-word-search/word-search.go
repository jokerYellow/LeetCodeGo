package leetcode

/*
https://leetcode.com/problems/word-search/
79. Word Search
Medium

2521

132

Favorite

Share
Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.
*/
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if _exist(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func _exist(board [][]byte, row, column int, word string, index int) bool {
	if index == len(word) {
		return true
	}
	if row < 0 || column < 0 || row >= len(board) || column >= len(board[0]) {
		return false
	}
	if board[row][column] != word[index] {
		return false
	}
	origin := board[row][column]
	board[row][column] = '#'
	if _exist(board, row-1, column, word, index+1) ||
		_exist(board, row+1, column, word, index+1) ||
		_exist(board, row, column-1, word, index+1) ||
		_exist(board, row, column+1, word, index+1) {
		return true
	}
	board[row][column] = origin
	return false
}

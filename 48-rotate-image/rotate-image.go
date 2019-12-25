package leetcode

/*
https://leetcode.com/problems/rotate-image/

48. Rotate Image
Medium

2176

186

Add to List

Share
You are given an n x n 2D matrix representing an image.

Rotate the image by 90 degrees (clockwise).

Note:

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

Example 1:

Given input matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

rotate the input matrix in-place such that it becomes:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
Example 2:

Given input matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

rotate the input matrix in-place such that it becomes:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
*/
func rotate(matrix [][]int) {
	_rotate(matrix, 0)
}

func _rotate(matrix [][]int, index int) {
	length := len(matrix)
	if index >= ((length + 1) / 2) {
		return
	}
	for i := index; i < length-index-1; i++ {
		p1x, p1y := index, i
		p2x, p2y := rotatePoint(p1x, p1y, length)
		p3x, p3y := rotatePoint(p2x, p2y, length)
		p4x, p4y := rotatePoint(p3x, p3y, length)
		matrix[p1x][p1y], matrix[p2x][p2y], matrix[p3x][p3y], matrix[p4x][p4y] = matrix[p4x][p4y], matrix[p1x][p1y], matrix[p2x][p2y], matrix[p3x][p3y]
	}
	_rotate(matrix, index+1)
}

func rotatePoint(x, y, length int) (int, int) {
	return y, length - 1 - x
}

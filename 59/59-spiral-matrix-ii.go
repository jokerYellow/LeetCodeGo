package _9

/*
https://leetcode.com/problems/spiral-matrix-ii/
59. Spiral Matrix II
Medium

462

82

Favorite

Share
Given a positive integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.

Example:

Input: 3
Output:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
*/

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	i := 0
	j := 0
	round := 0

	for index := 1; index <= n*n; index++ {
		topBounds := round
		rightBounds := n - round - 1
		bottomBounds := n - round - 1
		leftBounds := round
		if i == topBounds && j <= rightBounds {
			matrix[i][j] = index
			if j == rightBounds {
				i++
			} else {
				j++
			}
		} else if j == rightBounds && i <= bottomBounds {
			matrix[i][j] = index
			if i == bottomBounds {
				j--
			} else {
				i++
			}
		} else if i == bottomBounds && j >= leftBounds {
			matrix[i][j] = index
			if j == leftBounds {
				i--
			} else {
				j--
			}
		} else if j == leftBounds && i >= topBounds {
			matrix[i][j] = index
			if i == topBounds+1 {
				round++
				j++
			} else {
				i--
			}
		}
	}
	return matrix
}

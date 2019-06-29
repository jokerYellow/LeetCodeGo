package _4

/*
54. Spiral Matrix
Medium

1137

418

Favorite

Share
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

Example 1:

Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output: [1,2,3,6,9,8,7,4,5]
Example 2:

Input:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
*/

func spiralOrder(matrix [][]int) []int {
	rowsCount := len(matrix)
	if rowsCount == 0 {
		return []int{}
	}
	columnsCount := len(matrix[0])
	rt := make([]int, rowsCount*columnsCount)
	i := 0
	j := 0
	round := 0

	for index := 0; index < rowsCount*columnsCount; index++ {
		topBounds := round
		rightBounds := columnsCount - round - 1
		bottomBounds := rowsCount - round - 1
		leftBounds := round
		var value int
		if i == topBounds && j <= rightBounds {
			value = matrix[i][j]
			if j == rightBounds {
				i++
			} else {
				j++
			}
		} else if j == rightBounds && i <= bottomBounds {
			value = matrix[i][j]
			if i == bottomBounds {
				j--
			} else {
				i++
			}
		} else if i == bottomBounds && j >= leftBounds {
			value = matrix[i][j]
			if j == leftBounds {
				i--
			} else {
				j--
			}
		} else if j == leftBounds && i >= topBounds {
			value = matrix[i][j]
			if i == topBounds+1 {
				round++
				j++
			} else {
				i--
			}
		}
		rt[index] = value
	}
	return rt
}

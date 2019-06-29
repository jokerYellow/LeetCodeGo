package _85_spiral_matrix_iii

import "fmt"

/*
https://leetcode.com/problems/spiral-matrix-iii/

885. Spiral Matrix III
Medium

98

157

Favorite

Share
On a 2 dimensional grid with R rows and C columns, we start at (r0, c0) facing east.

Here, the north-west corner of the grid is at the first row and column, and the south-east corner of the grid is at the last row and column.

Now, we walk in a clockwise spiral shape to visit every position in this grid.

Whenever we would move outside the boundary of the grid, we continue our walk outside the grid (but may return to the grid boundary later.)

Eventually, we reach all R * C spaces of the grid.

Return a list of coordinates representing the positions of the grid in the order they were visited.



Example 1:

Input: R = 1, C = 4, r0 = 0, c0 = 0
Output: [[0,0],[0,1],[0,2],[0,3]]




Example 2:

Input: R = 5, C = 6, r0 = 1, c0 = 4
Output: [[1,4],[1,5],[2,5],[2,4],[2,3],[1,3],[0,3],[0,4],[0,5],[3,5],[3,4],[3,3],[3,2],[2,2],[1,2],[0,2],[4,5],[4,4],[4,3],[4,2],[4,1],[3,1],[2,1],[1,1],[0,1],[4,0],[3,0],[2,0],[1,0],[0,0]]




Note:

1 <= R <= 100
1 <= C <= 100
0 <= r0 < R
0 <= c0 < C
Accepted
9,243
Submissions
14,346
*/

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	corner := abs(R - r0)
	if corner < abs(C-c0) {
		corner = abs(C - c0)
	}
	matrix := make([][]int, R*C)
	for i := range matrix {
		matrix[i] = make([]int, 2)
	}
	r := r0
	c := c0
	round := 1
	index := 0
	if r < R && c < C && r >= 0 && c >= 0 {
		matrix[index] = []int{r, c}
		index++
		if index == R*C {
			return matrix
		}
	}

	for index < R*C {
		//1,1,2,2,3,3,4,4,5,5,6,6
		for i := 1; i <= round; i++ {
			c = c + 1
			fmt.Printf("r:%d,c:%d\n", r, c)
			if r < R && c < C && r >= 0 && c >= 0 {
				matrix[index] = []int{r, c}
				index++
				if index == R*C {
					return matrix
				}
			}
		}

		for i := 1; i <= round; i++ {
			r = r + 1
			fmt.Printf("r:%d,c:%d\n", r, c)
			if r < R && c < C && r >= 0 && c >= 0 {
				matrix[index] = []int{r, c}
				index++
				if index == R*C {
					return matrix
				}
			}
		}

		for i := 1; i <= round+1; i++ {
			c = c - 1
			fmt.Printf("r:%d,c:%d\n", r, c)
			if r < R && c < C && r >= 0 && c >= 0 {
				matrix[index] = []int{r, c}
				index++
				if index == R*C {
					return matrix
				}
			}
		}

		for i := 1; i <= round+1; i++ {
			r = r - 1
			fmt.Printf("r:%d,c:%d\n", r, c)
			if r < R && c < C && r >= 0 && c >= 0 {
				matrix[index] = []int{r, c}
				index++
				if index == R*C {
					return matrix
				}
			}
		}
		round += 2
	}
	return matrix
}

package _74

/*
https://leetcode.com/problems/dungeon-game/
174. Dungeon Game
Hard

609

14

Favorite

Share
The demons had captured the princess (P) and imprisoned her in the bottom-right corner of a dungeon. The dungeon consists of M x N rooms laid out in a 2D grid. Our valiant knight (K) was initially positioned in the top-left room and must fight his way through the dungeon to rescue the princess.

The knight has an initial health point represented by a positive integer. If at any point his health point drops to 0 or below, he dies immediately.

Some of the rooms are guarded by demons, so the knight loses health (negative integers) upon entering these rooms; other rooms are either empty (0's) or contain magic orbs that increase the knight's health (positive integers).

In order to reach the princess as quickly as possible, the knight decides to move only rightward or downward in each step.



Write a function to determine the knight's minimum initial health so that he is able to rescue the princess.

For example, given the dungeon below, the initial health of the knight must be at least 7 if he follows the optimal path RIGHT-> RIGHT -> DOWN -> DOWN.

-2 (K)	-3	   3
-5	    -10	   1
10 	    30	   -5 (P)

*/

func calculateMinimumHP(dungeon [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		} else {
			return b
		}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	row := len(dungeon)

	if row == 0 || dungeon == nil {
		return 0
	}

	column := len(dungeon[0])

	hp := make([][]int, row)
	for i := range hp {
		hp[i] = make([]int, column)
	}

	for i := row - 1; i >= 0; i-- {
		for j := column - 1; j >= 0; j-- {
			if i == row-1 && j == column-1 {
				hp[i][j] = max(1, 1-dungeon[i][j])
			} else if i == row-1 {
				hp[i][j] = max(1, hp[i][j+1]-dungeon[i][j])
			} else if j == column-1 {
				hp[i][j] = max(1, hp[i+1][j]-dungeon[i][j])
			} else {
				hp[i][j] = min(max(1, hp[i+1][j]-dungeon[i][j]), max(1, hp[i][j+1]-dungeon[i][j]))
			}
		}
	}

	return hp[0][0]
}

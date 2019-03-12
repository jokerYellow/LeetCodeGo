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

type hp struct {
	Sum int
	K   int
}

func calculateMinimumHP(dungeon [][]int) int {
	coloumn := len(dungeon)
	row := len(dungeon[0])
	rt := calculateMinimumHPDetail(dungeon, coloumn-1, row-1)
	min := -1
	for _, v := range rt {
		if min == -1 {
			min = v.K
			continue
		} else if min > v.K {
			min = v.K
		}
	}
	return min
}

func calculateMinimumHPDetail(dungeon [][]int, row, column int) []hp {
	if row == -1 || column == -1 {
		return []hp{hp{-1, -1}}
	}
	if row == 0 && column == 0 {
		p := dungeon[0][0]
		if p > 0 {
			return []hp{hp{1 + p, 1}}
		} else {
			return []hp{hp{1, 1 - p}}
		}
	}
	items := []hp{}
	tops := calculateMinimumHPDetail(dungeon, row-1, column)
	lefts := calculateMinimumHPDetail(dungeon, row, column-1)
	for _, t := range tops {
		items = append(items, t)
	}
	for _, t := range lefts {
		items = append(items, t)
	}

	p := dungeon[row][column]

	rt := []hp{}
	for _, t := range items {
		if t.K < 0 {
			continue
		}
		nk := 0
		if t.Sum+p > 0 {
			nk = t.K

		} else {
			nk = t.K - t.Sum - p + 1
		}
		rt = append(rt, hp{nk - t.K + t.Sum + p, nk})
	}
	return rt
}

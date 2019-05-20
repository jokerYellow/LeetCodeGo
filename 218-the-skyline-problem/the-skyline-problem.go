package _18_the_skyline_problem

import (
	"sort"
)

/*
https://leetcode.com/problems/the-skyline-problem/
218. The Skyline Problem
Hard

1138

62

Favorite

Share
A city's skyline is the outer contour of the silhouette formed by all the buildings in that city when viewed from a distance. Now suppose you are given the locations and height of all the buildings as shown on a cityscape photo (Figure A), write a program to output the skyline formed by these buildings collectively (Figure B).

Buildings  Skyline Contour
The geometric information of each building is represented by a triplet of integers [Li, Ri, Hi], where Li and Ri are the x coordinates of the left and right edge of the ith building, respectively, and Hi is its height. It is guaranteed that 0 ≤ Li, Ri ≤ INT_MAX, 0 < Hi ≤ INT_MAX, and Ri - Li > 0. You may assume all buildings are perfect rectangles grounded on an absolutely flat surface at height 0.

For instance, the dimensions of all buildings in Figure A are recorded as: [ [2 9 10], [3 7 15], [5 12 12], [15 20 10], [19 24 8] ] .

The output is a list of "key points" (red dots in Figure B) in the format of [ [x1,y1], [x2, y2], [x3, y3], ... ] that uniquely defines a skyline. A key point is the left endpoint of a horizontal line segment. Note that the last key point, where the rightmost building ends, is merely used to mark the termination of the skyline, and always has zero height. Also, the ground in between any two adjacent buildings should be considered part of the skyline contour.

For instance, the skyline in Figure B should be represented as:[ [2 10], [3 15], [7 12], [12 0], [15 10], [20 8], [24, 0] ].

Notes:

The number of buildings in any input list is guaranteed to be in the range [0, 10000].
The input list is already sorted in ascending order by the left x position Li.
The output list must be sorted by the x position.
There must be no consecutive horizontal lines of equal height in the output skyline. For instance, [...[2 3], [4 5], [7 5], [11 5], [12 7]...] is not acceptable; the three lines of height 5 should be merged into one in the final output as such: [...[2 3], [4 5], [12 7], ...]
*/
type Item struct {
	isLeft bool
	index  int
	height int
}

// TODO: refactor
func getSkyline(buildings [][]int) [][]int {
	var removeRepeated [][]int
	var lastOne []int
	for _, v := range buildings {
		if len(lastOne) > 0 && lastOne[0] == v[0] && lastOne[1] == v[1] && lastOne[2] == v[2] {
			continue
		} else {
			removeRepeated = append(removeRepeated, v)
			lastOne = v
		}
	}
	buildings = removeRepeated

	var rt [][]int
	items := make([]*Item, len(buildings)*2)
	for i, v := range buildings {
		left := Item{isLeft: true, index: v[0], height: v[2]}
		items[2*i] = &left
		right := Item{isLeft: false, index: v[1], height: v[2]}
		items[2*i+1] = &right
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].index < items[j].index
	})

	for _, item := range items {
		if item.isLeft {
			lowwerThanSomeOne := false
			for _, v := range buildings {
				if v[0] <= item.index && v[1] >= item.index {
					if v[2] > item.height ||
						(v[2] == item.height && item.index == v[1]) ||
						(v[2] >= item.height && item.index > v[0] && item.index < v[1]) {
						lowwerThanSomeOne = true
						break
					}
				}
			}
			if lowwerThanSomeOne == false {
				rt = append(rt, []int{item.index, item.height})
			}
		} else {
			lowwerHeight := 0
			for _, v := range buildings {
				if v[0] <= item.index && v[1] >= item.index {
					if v[2] > item.height ||
						(v[2] == item.height && item.index == v[0]) ||
						(v[2] >= item.height && item.index > v[0] && item.index < v[1]) {
						lowwerHeight = -1
						break
					} else if v[2] < item.height {
						if v[1] > item.index {
							if lowwerHeight == 0 {
								lowwerHeight = v[2]
							} else if lowwerHeight < v[2] {
								lowwerHeight = v[2]
							}
						}
					}
				}
			}
			if lowwerHeight >= 0 {
				rt = append(rt, []int{item.index, lowwerHeight})
			}
		}

	}

	return rt
}

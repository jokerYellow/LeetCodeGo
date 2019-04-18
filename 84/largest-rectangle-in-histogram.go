package leetcode

/*
https://leetcode.com/problems/largest-rectangle-in-histogram/
84. Largest Rectangle in Histogram
Hard

1767

45

Favorite

Share
Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.




Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].




The largest rectangle is shown in the shaded area, which has area = 10 unit.



Example:

Input: [2,1,5,6,2,3]
Output: 10

*/

func largestRectangleArea(heights []int) int {
	return largestRectangleAreaReturnWithWidth(heights, 0)
}

func largestRectangleAreaReturnWithWidth(heights []int, height int) int {
	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0] + height
	}
	min := heights[0]
	for _, v := range heights {
		if min > v {
			min = v
		}
	}
	var areas [][]int
	var t []int
	for _, v := range heights {
		if v == min {
			if len(t) > 0 {
				areas = append(areas, t)
				t = []int{}
			}
		} else {
			t = append(t, v-min)
		}
	}
	if len(t) > 0 {
		areas = append(areas, t)
	}
	if len(areas) == 0 {
		return (min + height) * len(heights)
	}
	max := (min + height) * len(heights)
	for _, v := range areas {
		tmp := largestRectangleAreaReturnWithWidth(v, height+min)
		if tmp > max {
			max = tmp
		}
	}
	return max
}

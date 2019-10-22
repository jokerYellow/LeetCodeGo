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

	max := func(l, r int) int {
		if l > r {
			return l
		}
		return r
	}

	length := len(heights)
	lessLeft := make([]int, length)
	lessRight := make([]int, length)

	for i := 0; i < length; i++ {
		p := i
		for p >= 0 && heights[p] >= heights[i] {
			p = p - 1
		}
		lessLeft[i] = p + 1
	}

	for i := length - 1; i >= 0; i-- {
		p := i
		for p <= (length-1) && heights[p] >= heights[i] {
			p = p + 1
		}
		lessRight[i] = p - 1
	}

	maximum := 0

	for i := 0; i < length; i++ {
		maximum = max(maximum, (lessRight[i]-lessLeft[i]+1)*heights[i])
	}

	return maximum
}

func largestRectangleArea1(heights []int) int {
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

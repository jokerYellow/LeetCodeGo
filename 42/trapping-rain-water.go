package leetcode

/*
https://leetcode.com/problems/trapping-rain-water/
42. Trapping Rain Water
Hard

3409

61

Favorite

Share
Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it is able to trap after raining.


The above elevation map is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped. Thanks Marcos for contributing this image!

Example:
0 1 0 0 1 0
  0 1 1 0

Input: [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6

*/

func trap(height []int) int {

	return trap3(height)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//brute force
//time complexity O(n2)
//space complexity O(1)
func trap1(height []int) int {
	traps := 0
	length := len(height)
	for i, v := range height {
		var leftMax, rightMax int
		for left := i; left >= 0; left-- {
			leftMax = max(height[left], leftMax)
		}
		for right := i; right < length; right++ {
			rightMax = max(height[right], rightMax)
		}
		traps += min(leftMax, rightMax) - v
	}
	return traps
}

//store the heights of each one
//time complexity O(n)
//space complexity O(n)
func trap2(height []int) int {

	length := len(height)
	maxLeft := make([]int, length)
	if length == 0 {
		return 0
	}

	traps := 0
	for i := 1; i < length; i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i])
	}

	maxRight := make([]int, length)
	maxRight[length-1] = height[length-1]

	for j := length - 2; j >= 0; j-- {
		maxRight[j] = max(maxRight[j+1], height[j])
	}

	for i := 0; i < length; i++ {
		traps += min(maxRight[i], maxLeft[i]) - height[i]
	}

	return traps
}

//#TODO: make me clear
func trap3(height []int) int {

	length := len(height)
	if length == 0 {
		return 0
	}
	leftIndex := 0
	rightIndex := length - 1

	maxLeft := height[leftIndex]
	maxRight := height[rightIndex]

	traps := 0
	for leftIndex < rightIndex {
		for maxLeft <= maxRight && leftIndex < rightIndex {
			leftIndex += 1
			value := height[leftIndex]
			maxLeft = max(maxLeft, value)
			if maxLeft <= maxRight {
				traps += min(maxLeft, maxRight) - value
			}else{
				break
			}
		}
		for maxRight <= maxLeft && leftIndex < rightIndex {
			rightIndex -= 1
			value := height[rightIndex]
			maxRight = max(maxRight, value)
			if maxLeft >= maxRight{
				traps += min(maxRight, maxLeft) - value
			}else{
				break
			}
		}
	}

	return traps
}

package leetcode

func maxArea(height []int) int {
	max := 0
	left := 0
	right := len(height) - 1
	for left < right {
		area := calculateArea(left, right, height)
		if area > max {
			max = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func calculateArea(left, right int, height []int) int {
	lowBounds := height[left]
	if height[left] > height[right] {
		lowBounds = height[right]
	}
	return (right - left) * lowBounds
}

package binarySearch

func binarySearch(nums []int, value int) int {
	if len(nums) == 0 {
		return -1
	}
	return _binarySearch(nums, value, 0, len(nums)-1)
}

func _binarySearch(nums []int, value int, start, end int) int {
	if nums[start] > value || nums[end] < value {
		return -1
	}
	mid := start + (end-start)/2
	if nums[mid] > value {
		return _binarySearch(nums, value, start, mid-1)
	} else if nums[mid] < value {
		return _binarySearch(nums, value, mid+1, end)
	} else {
		return mid
	}
}

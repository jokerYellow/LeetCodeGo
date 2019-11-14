package midSearch

func midSearch(nums []int, value int) int {
	return midSearchSub(nums, value, 0, len(nums)-1)
}

func midSearchSub(nums []int, value int, start, end int) int {
	if start > end {
		return -1
	}
	mid := start + (end-start)/2
	if nums[mid] > value {
		return midSearchSub(nums, value, start, mid-1)
	} else if nums[mid] < value {
		return midSearchSub(nums, value, mid+1, end)
	}
	return mid
}

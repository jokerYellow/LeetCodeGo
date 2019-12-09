package leetcode

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 || target < nums[0] || target > nums[len(nums)-1] {
		return result
	}
	left := binarySearch(nums, target, 0, len(nums)-1, false)
	right := binarySearch(nums, target, 0, len(nums)-1, true)
	return []int{left, right}
}

func binarySearch(nums []int, target, left, right int, bigger bool) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if nums[left] > target || nums[right] < target || len(nums) == 0 {
		return -1
	}
	if target == nums[mid] && (mid == 0 || nums[mid-1] < nums[mid]) && !bigger {
		return mid
	} else if target == nums[mid] && (mid == len(nums)-1 || nums[mid+1] > nums[mid]) && bigger {
		return mid
	}
	if bigger {
		if nums[mid] <= target {
			return binarySearch(nums, target, mid+1, right, bigger)
		} else {
			return binarySearch(nums, target, left, mid-1, bigger)
		}
	} else {
		if nums[mid] < target {
			return binarySearch(nums, target, mid+1, right, bigger)
		} else {
			return binarySearch(nums, target, left, mid-1, bigger)
		}
	}
}

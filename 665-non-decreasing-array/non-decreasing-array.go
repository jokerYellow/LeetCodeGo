package _65_non_decreasing_array

func checkPossibility(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		v,next := nums[i],nums[i+1]
		if v > next {
			nums[i] = next
			nums[i+1] = next
			if !isSort(nums) {
				nums[i] = v
				nums[i+1] = v
				if !isSort(nums) {
					return false
				}
			}
		}
	}
	return true
}

func isSort(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
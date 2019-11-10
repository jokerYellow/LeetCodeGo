package seperate

func seperate(nums []int) []int {
	l := 0
	r := len(nums) - 1
	for l < r {
		for nums[l]%2 == 1 && l < r {
			l++
		}
		for nums[r]%2 == 0 && l < r {
			r--
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	return nums
}

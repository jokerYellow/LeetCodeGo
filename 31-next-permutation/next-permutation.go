package leetcode

func nextPermutation(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	lastIndex := len(nums) - 1
	lastValue := nums[lastIndex]
	for i := len(nums) - 2; i >= 0; i-- {
		if lastValue > nums[i] {
			lastValue = nums[i]
			lastIndex = i
			break
		}
		lastValue = nums[i]
	}
	if lastIndex == len(nums)-1 {
		i := 0
		j := len(nums) - 1
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	} else {
		minBiggerIndex := lastIndex + 1
		for i := lastIndex + 2; i < len(nums); i++ {
			if nums[i] > lastValue {
				minBiggerIndex = i
			} else if nums[i] < lastValue {
				break
			}
		}
		nums[lastIndex], nums[minBiggerIndex] = nums[minBiggerIndex], nums[lastIndex]
		i := lastIndex + 1
		j := len(nums) - 1
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	return nums
}

package two_sum_ii_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	for {
		s := numbers[l] + numbers[r]
		if s > target {
			r--
		} else if s == target {
			break
		} else {
			l++
		}
	}
	return []int{l + 1, r + 1}
}

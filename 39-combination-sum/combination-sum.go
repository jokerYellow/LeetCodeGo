package leetcode

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	if len(candidates) == 0 {
		return result
	}
	if candidates[0] != 0 && target%candidates[0] == 0 {
		count := target / candidates[0]
		rt := make([]int, count)
		for i := 0; i < count; i++ {
			rt[i] = candidates[0]
		}
		result = append(result, rt)
	}
	if len(candidates) > 1 {
		left := candidates[0]
		lefts := []int{left}
		for left < target {
			for _, item := range combinationSum(candidates[1:], target-left) {
				valid := []int{}
				valid = append(valid, lefts...)
				valid = append(valid, item...)
				result = append(result, valid)
			}
			left += candidates[0]
			lefts = append(lefts, candidates[0])
		}
		result = append(result, combinationSum(candidates[1:], target)...)
	}
	return result
}

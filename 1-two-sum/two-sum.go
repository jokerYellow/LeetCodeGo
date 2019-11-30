package leetcode

func twoSum(nums []int, target int) []int {
	numsMap := map[int]int{}
	for index, value := range nums {
		if preIndex, ok := numsMap[target-value]; ok {
			return []int{preIndex, index}
		}
		numsMap[value] = index
	}
	return []int{}
}

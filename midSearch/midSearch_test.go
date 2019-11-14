package midSearch

import "testing"

func Test(t *testing.T) {
	test(t, []int{1}, 2, -1)
	test(t, []int{}, 2, -1)
	test(t, []int{2}, 2, 0)

	test(t, []int{1, 2, 3, 4}, 2, 1)
}

func test(t *testing.T, nums []int, value int, expectIndex int) {
	if expectIndex != midSearch(nums, value) {
		t.Fail()
	}
}

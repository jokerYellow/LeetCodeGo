package midSearch

import (
	"fmt"
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

func Test(t *testing.T) {
	test(t, []int{1}, 2, -1)
	test(t, []int{}, 2, -1)
	test(t, []int{2}, 2, 0)

	test(t, []int{1, 2, 3, 4}, 2, 1)
	test(t, []int{1, 2, 3, 4, 5, 9}, 6, -1)
}

func test(t *testing.T, nums []int, value int, expectIndex int) {
	output := midSearch(nums, value)
	if expectIndex != output {
		fmt.Println(nums, value)
		utils.Print(expectIndex, output)
		t.Fail()
	}
}

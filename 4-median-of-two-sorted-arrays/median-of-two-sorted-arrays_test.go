package __median_of_two_sorted_arrays

import (
	"github.com/jokerYellow/leetcode/utils"
	"testing"
)

type testCase struct {
	nums1, nums2 []int
	expect       float64
}

func Test(t *testing.T) {
	cases := []testCase{
		{[]int{1, 3}, []int{2, 4}, 2.5},
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 3}, []int{}, 2},
	}
	for _, c := range cases {
		output := findMedianSortedArrays(c.nums1, c.nums2)
		if output != c.expect {
			utils.Print(c.expect, output)
			t.Fail()
		}
	}
}

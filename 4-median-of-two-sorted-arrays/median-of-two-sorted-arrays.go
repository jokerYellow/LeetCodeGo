package __median_of_two_sorted_arrays

/*
4. Median of Two Sorted Arrays
Hard

5397

796

Favorite

Share
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1Length := len(nums1)
	nums2Length := len(nums2)
	if nums1Length > nums2Length {
		return findMedianSortedArrays(nums2, nums1)
	}
	low := 0
	high := nums1Length
	k := (nums1Length + nums2Length + 1) >> 1
	midNums1 := 0
	midNums2 := 0
	for low <= high {
		midNums1 = low + (high-low)/2
		midNums2 = k - midNums1
		if midNums1 > 0 && nums1[midNums1-1] > nums2[midNums2] {
			high = midNums1 - 1
		} else if midNums1 < nums1Length && nums1[midNums1] < nums2[midNums2-1] {
			low = midNums1 + 1
		} else {
			break
		}
	}
	if (nums1Length+nums2Length)&1 == 1 {
		if nums1Length == 0 || midNums1 == 0 {
			return float64(nums2[midNums2-1])
		}
		return float64(max(nums2[midNums2-1], nums1[midNums1-1]))
	}
	maxLeft := 0
	minRight := 0
	if midNums1 == 0 {
		maxLeft = nums2[midNums2-1]
	} else if midNums2 == 0 {
		maxLeft = nums1[midNums1-1]
	} else {
		maxLeft = max(nums1[midNums1-1], nums2[midNums2-1])
	}

	if midNums1 == nums1Length {
		minRight = nums2[midNums2]
	} else if midNums2 == nums2Length {
		minRight = nums1[midNums1]
	} else {
		minRight = min(nums1[midNums1], nums2[midNums2])
	}
	return float64(maxLeft+minRight) / 2

}

func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

func min(l, r int) int {
	if l > r {
		return r
	}
	return l
}

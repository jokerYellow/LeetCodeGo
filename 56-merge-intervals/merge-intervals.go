package leetcode

import (
	"sort"
)

/*
https://leetcode.com/problems/merge-intervals/
56. Merge Intervals
Medium

2977

231

Add to List

Share
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.

Accepted
465,663
Submissions
1,242,212
*/
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || (intervals[i][0] == intervals[j][0] && intervals[i][1] <= intervals[j][1])
	})
	var result [][]int
	curr := intervals[0]
	for i := 1; i < len(intervals); i++ {
		t := intervals[i]
		if t[0] <= curr[1] && t[1] > curr[1] {
			curr[1] = t[1]
		} else if t[0] > curr[1] {
			result = append(result, curr)
			curr = t
		}
	}
	if len(curr) == 2 {
		result = append(result, curr)
	}
	return result
}

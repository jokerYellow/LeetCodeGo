package leetcode

/*
https://leetcode.com/problems/insert-interval/
57. Insert Interval
Hard

714

89

Favorite

Share
Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
Example 2:

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
*/

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	var rt []Interval
	var currentInterval *Interval
	if len(intervals) == 0 || intervals == nil {
		rt = append(rt, newInterval)
		return rt
	}

	for i, v := range intervals {
		if currentInterval == nil && newInterval.End < v.Start {
			if i > 0 {
				if intervals[i-1].End < newInterval.Start {
					rt = append(rt, newInterval)
				}
			} else if newInterval.End < v.Start {
				rt = append(rt, newInterval)
			}
			rt = append(rt, v)
		} else if currentInterval == nil && newInterval.Start > v.End {
			rt = append(rt, v)
			if i == len(intervals)-1 {
				rt = append(rt, newInterval)
			}
		} else if currentInterval == nil && newInterval.Start <= v.End {
			if newInterval.End < v.Start {
				rt = append(rt, newInterval)
				rt = append(rt, v)
				continue
			}
			currentInterval = new(Interval)
			currentInterval.Start = min(newInterval.Start, v.Start)
			if newInterval.End <= v.End {
				currentInterval.End = v.End
				rt = append(rt, *currentInterval)
				currentInterval = nil
			} else {
				currentInterval.End = newInterval.End
			}
		} else if currentInterval != nil && currentInterval.End >= v.Start && v.End > currentInterval.End {
			currentInterval.End = v.End
			rt = append(rt, *currentInterval)
			currentInterval = nil
		} else if currentInterval != nil && currentInterval.End < v.Start {
			rt = append(rt, *currentInterval)
			currentInterval = nil
			rt = append(rt, v)
		}
	}
	if currentInterval != nil {
		rt = append(rt, *currentInterval)
	}
	return rt
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

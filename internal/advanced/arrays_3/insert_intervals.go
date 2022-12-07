package arrays3

import "github.com/chinmayb13/schrodinger-cat/internal/helper"

/*
Problem Description
Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).
You may assume that the intervals were initially sorted according to their start times.

Problem Constraints
0 <= |intervals| <= 105

Input Format
First argument is the vector of intervals
second argument is the new interval to be merged

Output Format
Return the vector of intervals after merging

Example Input
Input 1:
Given intervals [1, 3], [6, 9] insert and merge [2, 5] .

Input 2:
Given intervals [1, 3], [6, 9] insert and merge [2, 6] .

Example Output
Output 1:
[ [1, 5], [6, 9] ]

Output 2:
[ [1, 9] ]

Example Explanation
Explanation 1:
(2,5) does not completely merge the given intervals

Explanation 2:
(2,6) completely merges the given intervals
*/
type Interval struct {
	start, end int
}

func InsertIntervalsII(intervals []Interval, newInterval Interval) []Interval {
	var result []Interval

	for i := 0; i < len(intervals); i++ {
		if newInterval.end < intervals[i].start {
			result = append(result, newInterval)
			for ; i < len(intervals); i++ {
				result = append(result, intervals[i])
			}
			return result
		} else if intervals[i].end < newInterval.start {
			result = append(result, intervals[i])
		} else {
			newInterval.start = helper.Min(newInterval.start, intervals[i].start)
			newInterval.end = helper.Max(newInterval.end, intervals[i].end)
		}
	}
	result = append(result, newInterval)
	return result
}

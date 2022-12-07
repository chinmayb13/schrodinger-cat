package arrays3

import (
	"sort"

	"github.com/chinmayb13/schrodinger-cat/internal/helper"
)

/*
Problem Description
Given a collection of intervals, merge all overlapping intervals.

Problem Constraints
1 <= Total number of intervals <= 100000.

Input Format
First argument is a list of intervals.

Output Format
Return the sorted list of intervals after merging all the overlapping intervals.

Example Input
Input 1:
[1,3],[2,6],[8,10],[15,18]

Example Output
Output 1:
[1,6],[8,10],[15,18]

Example Explanation
Explanation 1:
Merge intervals [1,3] and [2,6] -> [1,6].
so, the required answer after merging is [1,6],[8,10],[15,18].
No more overlapping intervals present.
*/
func MergeIntervals(A []Interval) []Interval {
	sort.Slice(A, func(i, j int) bool {
		return (A[i].start < A[j].start) || ((A[i].start == A[j].start) && (A[i].end < A[j].end))
	})
	var result []Interval
	//initialise l and r with first interval
	l := A[0].start
	r := A[0].end
	for i := 1; i < len(A); i++ {
		//if current interval's start overlaps with R
		if A[i].start <= r {
			r = helper.Max(r, A[i].end)
		} else {
			//if current interval doesn't overlap
			result = append(result, Interval{
				start: l,
				end:   r,
			})
			l = A[i].start
			r = A[i].end
		}
	}
	result = append(result, Interval{
		start: l,
		end:   r,
	})
	return result
}

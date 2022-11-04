package sorting

import "sort"

func CanAttendMeetings(intervals [][]int) bool {
	if len(intervals) <= 1 {
		return true
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}

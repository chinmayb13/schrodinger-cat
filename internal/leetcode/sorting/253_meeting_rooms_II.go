package sorting

import "sort"

func MinMeetingRooms(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 1
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	bookedRoom := make(map[int]bool)

	count := 1
	for i := 1; i < len(intervals); i++ {
		sameRoom := false
		for j := 0; j < i; j++ {
			if (intervals[i][0] >= intervals[j][1]) && !bookedRoom[j] {
				sameRoom = true
				bookedRoom[j] = true
				break
			}
		}

		if !sameRoom {
			count++
		}
	}
	return count
}

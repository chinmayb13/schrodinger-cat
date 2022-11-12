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

func MinMeetingRoomsAlt(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 1
	}
	maxInterval := -1
	for i := range intervals {
		if intervals[i][1] > maxInterval {
			maxInterval = intervals[i][1]
		}
	}

	timeslot := make([]int, maxInterval+1)
	for i := range intervals {
		timeslot[intervals[i][0]]++
		timeslot[intervals[i][1]]--
	}

	minRooms := 0
	for i := 1; i < len(timeslot); i++ {
		timeslot[i] += timeslot[i-1]
		if timeslot[i] > minRooms {
			minRooms = timeslot[i]
		}
	}
	return minRooms
}

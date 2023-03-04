package heaps1

import "math"

func getLeftMinChild(idx int, inp []int) int {
	childIdx := (idx << 1) + 1
	if childIdx >= len(inp) {
		return math.MaxInt
	}
	return inp[childIdx]
}

func getRightMinChild(idx int, inp []int) int {
	childIdx := (idx << 1) + 2
	if childIdx >= len(inp) {
		return math.MaxInt
	}
	return inp[childIdx]
}

func min(args ...int) int {
	ans := math.MaxInt
	for i := range args {
		if args[i] < ans {
			ans = args[i]
		}
	}
	return ans
}

func minHeapify(idx int, inp []int) {
	if len(inp) < 2 {
		return
	}
	for {
		lc := getLeftMinChild(idx, inp)
		rc := getRightMinChild(idx, inp)
		val := inp[idx]
		minVal := min(val, lc, rc)
		if minVal == val {
			return
		} else if minVal == lc {
			inp[idx], inp[(idx<<1)+1] = lc, val
			idx = (idx << 1) + 1
		} else {
			inp[idx], inp[(idx<<1)+2] = rc, val
			idx = (idx << 1) + 2
		}
	}
}

func createMinHeap(inp []int) {
	for i := (len(inp) - 2) >> 1; i >= 0; i-- {
		minHeapify(i, inp)
	}
}

func createMaxHeap(inp []int) {
	for i := (len(inp) - 2) >> 1; i >= 0; i-- {
		maxHeapify(i, inp)
	}
}

func maxHeapify(idx int, inp []int) {
	if len(inp) < 2 {
		return
	}
	for {
		lc := getMaxLC(idx, inp)
		rc := getMaxRC(idx, inp)
		val := inp[idx]
		maxVal := maxElement(val, lc, rc)
		if maxVal == val {
			return
		} else if maxVal == lc {
			inp[idx], inp[(idx<<1)+1] = lc, val
			idx = (idx << 1) + 1
		} else {
			inp[idx], inp[(idx<<1)+2] = rc, val
			idx = (idx << 1) + 2
		}
	}
}

func getMaxLC(idx int, inp []int) int {
	childIdx := (idx << 1) + 1
	if childIdx >= len(inp) {
		return -1
	}
	return inp[childIdx]
}

func getMaxRC(idx int, inp []int) int {
	childIdx := (idx << 1) + 2
	if childIdx >= len(inp) {
		return -1
	}
	return inp[childIdx]
}

func maxElement(args ...int) int {
	ans := 0
	for i := range args {
		if args[i] > ans {
			ans = args[i]
		}
	}
	return ans
}

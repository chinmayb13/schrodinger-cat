package heaps2

type minIntHeap []int

func (h minIntHeap) Len() int           { return len(h) }
func (h minIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minIntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *minIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type maxIntHeap []int

func (h maxIntHeap) Len() int           { return len(h) }
func (h maxIntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxIntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *maxIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type info struct {
	idx1  int
	idx2  int
	value int
}

type maxStructHeap []*info

func (h maxStructHeap) Len() int           { return len(h) }
func (h maxStructHeap) Less(i, j int) bool { return h[i].value > h[j].value }
func (h maxStructHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxStructHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*info))
}

func (h *maxStructHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return x
}

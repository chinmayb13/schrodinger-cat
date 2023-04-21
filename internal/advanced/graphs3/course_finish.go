package graphs3

import "container/heap"

func CanCompleteCourses(A int, B []int, C []int) int {
	var ans []int
	inDegree := make([]int, A+1)
	adjList := make([][]int, A+1)
	//create adjacency list
	for i := 0; i < len(B); i++ {
		inDegree[C[i]]++
		adjList[B[i]] = append(adjList[B[i]], C[i])
	}

	//make minheap
	mh := make(minIntHeap, 0)
	heap.Init(&mh)
	for i := 1; i <= A; i++ {
		//insert all zero indegree vertices to minHeap
		if inDegree[i] == 0 {
			heap.Push(&mh, i)
		}
	}

	//for each minHeap item, delete its outgoing edges
	//by decreasing the indegree of the vertices present in the adjacency list of the item
	for len(mh) > 0 {
		u := mh[0]
		heap.Pop(&mh)
		ans = append(ans, u)
		for j := range adjList[u] {
			v := adjList[u][j]
			inDegree[v]--
			//if indegree is 0, push to minHeap
			if inDegree[v] == 0 {
				heap.Push(&mh, v)
			}
		}
	}
	if len(ans) == A {
		return 1
	}
	return 0
}

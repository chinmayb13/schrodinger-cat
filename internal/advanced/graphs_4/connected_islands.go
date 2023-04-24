package graphs4

import "container/heap"

/*
Problem Description
There are A islands and there are M bridges connecting them. Each bridge has some cost attached to it.
We need to find bridges with minimal cost such that all islands are connected.
It is guaranteed that input data will contain at least one possible scenario in which all islands are connected with each other.

Problem Constraints
1 <= A, M <= 6*104
1 <= B[i][0], B[i][1] <= A
1 <= B[i][2] <= 103

Input Format
The first argument contains an integer, A, representing the number of islands.
The second argument contains an 2-d integer matrix, B, of size M x 3 where Island B[i][0] and B[i][1] are connected using a bridge of cost B[i][2].

Output Format
Return an integer representing the minimal cost required.

Example Input
Input 1:
A = 4
B = [  [1, 2, 1]
       [2, 3, 4]
       [1, 4, 3]
       [4, 3, 2]
       [1, 3, 10]  ]

Input 2:
A = 4
B = [  [1, 2, 1]
       [2, 3, 2]
       [3, 4, 4]
       [1, 4, 3]   ]

Example Output
Output 1:
6

Output 2:
6

Example Explanation
Explanation 1:
We can choose bridges (1, 2, 1), (1, 4, 3) and (4, 3, 2), where the total cost incurred will be (1 + 3 + 2) = 6.

Explanation 2:
We can choose bridges (1, 2, 1), (2, 3, 2) and (1, 4, 3), where the total cost incurred will be (1 + 2 + 3) = 6.
*/
func GetMinimumCost(A int, B [][]int) int {
	adjList := make([][]pair, A+1)
	//create adj list
	for i := range B {
		adjList[B[i][0]] = append(adjList[B[i][0]], pair{dest: B[i][1], cost: B[i][2]})
		adjList[B[i][1]] = append(adjList[B[i][1]], pair{dest: B[i][0], cost: B[i][2]})
	}

	//create visited array
	visited := make([]bool, A+1)
	mh := make(minTripletHeap, 0)
	heap.Init(&mh)

	//consider 1 as the source and insert the edges in the min Heap
	for i := range adjList[1] {
		p := adjList[1][i]
		heap.Push(&mh, triplet{src: 1, dest: p.dest, cost: p.cost})
	}

	visited[1] = true
	ans := 0
	for len(mh) > 0 {
		tplt := mh[0]
		heap.Pop(&mh)
		//if the edge is already considered, skip
		if visited[tplt.src] && visited[tplt.dest] {
			continue
		}

		//add the dest of the edge as visted
		visited[tplt.dest] = true
		ans += tplt.cost

		//iterate through the outgoing edges of the dest and if unvisited, insert into minHeap
		for i := range adjList[tplt.dest] {
			p := adjList[tplt.dest][i]
			if !visited[p.dest] {
				heap.Push(&mh, triplet{src: tplt.dest, dest: p.dest, cost: p.cost})
			}
		}
	}
	return ans

}

type triplet struct {
	src  int
	dest int
	cost int
}

type pair struct {
	dest int
	cost int
}

type minTripletHeap []triplet

func (h minTripletHeap) Len() int           { return len(h) }
func (h minTripletHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h minTripletHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minTripletHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(triplet))
}

func (h *minTripletHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

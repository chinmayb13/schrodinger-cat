package graphs5

import "container/heap"

/*
Problem Description
Given a weighted undirected graph having A nodes and M weighted edges, and a source node C.
You have to find an integer array D of size A such that:
=> D[i] : Shortest distance form the C node to node i.
=> If node i is not reachable from C then -1.

Note:

There are no self-loops in the graph.
No multiple edges between two pair of vertices.
The graph may or may not be connected.
Nodes are numbered from 0 to A-1.
Your solution will run on multiple testcases. If you are using global variables make sure to clear them.

Problem Constraints
1 <= A <= 1e5
0 <= B[i][0],B[i][1] < A
0 <= B[i][2] <= 1e3
0 <= C < A

Input Format
The first argument given is an integer A, representing the number of nodes.
The second argument given is the matrix B of size M x 3, where nodes B[i][0] and B[i][1] are connected with an edge of weight B[i][2].
The third argument given is an integer C.

Output Format
Return the integer array D.

Example Input
Input 1:
A = 6
B = [   [0, 4, 9]

	[3, 4, 6]
	[1, 2, 1]
	[2, 5, 1]
	[2, 4, 5]
	[0, 3, 7]
	[0, 1, 1]
	[4, 5, 7]
	[0, 5, 1] ]

C = 4

Input 2:
A = 5
B = [   [0, 3, 4]

	[2, 3, 3]
	[0, 1, 9]
	[3, 4, 10]
	[1, 3, 8]  ]

C = 4

Example Output
Output 1:
D = [7, 6, 5, 6, 0, 6]

Output 2:
D = [14, 18, 13, 10, 0]

Example Explanation
Explanation 1:
All Paths can be considered from the node C to get shortest path

Explanation 2:
All Paths can be considered from the node C to get shortest path
*/
func FindShortesDistance(A int, B [][]int, C int) []int {
	adjList := make([][]nodePair, A)

	//make adjList
	for i := range B {
		u := B[i][0]
		v := B[i][1]
		w := B[i][2]
		adjList[u] = append(adjList[u], nodePair{node: v, weight: w})
		adjList[v] = append(adjList[v], nodePair{node: u, weight: w})
	}

	//make distance array for saving the distance from source
	dist := make([]int, A)
	for i := range dist {
		dist[i] = -1
	}

	mh := make(minHeap, 0)
	heap.Init(&mh)
	heap.Push(&mh, nodePair{node: C, weight: 0})

	for mh.Len() > 0 {
		pair := mh[0]
		vert := pair.node
		wgt := pair.weight
		heap.Pop(&mh)

		//if already visited, skip the node
		if dist[vert] != -1 {
			continue
		}

		//save the weight
		dist[vert] = wgt

		//iterate adjList
		for i := range adjList[vert] {
			p := adjList[vert][i]
			v := p.node
			w := p.weight

			//if not visited, push the minHeap
			if dist[v] == -1 {
				heap.Push(&mh, nodePair{node: v, weight: wgt + w})
			}
		}
	}
	return dist
}

type nodePair struct {
	node   int
	weight int
}

type minHeap []nodePair

func (mh minHeap) Len() int           { return len(mh) }
func (mh minHeap) Less(i, j int) bool { return mh[i].weight < mh[j].weight }
func (mh minHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }

func (mh *minHeap) Push(x interface{}) {
	*mh = append(*mh, x.(nodePair))
}

func (mh *minHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[:n-1]
	return x
}

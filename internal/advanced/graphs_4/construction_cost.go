package graphs4

import "container/heap"

/*
Problem Description
Given a graph with A nodes and C weighted edges. Cost of constructing the graph is the sum of weights of all the edges in the graph.
Find the minimum cost of constructing the graph by selecting some given edges such that we can reach every other node from the 1st node.

NOTE: Return the answer modulo 109+7 as the answer can be large.

Problem Constraints
1 <= A <= 100000
0 <= C <= 100000
1 <= B[i][0], B[i][1] <= N
1 <= B[i][2] <= 109

Input Format
First argument is an integer A.
Second argument is a 2-D integer array B of size C*3 denoting edges. B[i][0] and B[i][1] are connected by ith edge with weight B[i][2]

Output Format
Return an integer denoting the minimum construction cost.

Example Input
Input 1:
A = 3
B = [   [1, 2, 14]
        [2, 3, 7]
        [3, 1, 2]   ]
Input 2:

A = 3
B = [   [1, 2, 20]
        [2, 3, 17]  ]

Example Output
Output 1:
9

Output 2:
37

Example Explanation
Explanation 1:
We can take only two edges (2 -> 3 and 3 -> 1) to construct the graph.
we can reach the 1st node from 2nd and 3rd node using only these two edges.
So, the total cost of costruction is 9.

Explanation 2:
We have to take both the given edges so that we can reach the 1st node from 2nd and 3rd node.
*/

func GetMinConstructionCost(A int, B [][]int) int {
	var mod int = 1e9 + 7
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
		ans = (ans + tplt.cost) % mod

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

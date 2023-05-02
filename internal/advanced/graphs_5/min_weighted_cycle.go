package graphs5

import (
	"container/heap"
	"math"
)

/*
Problem Description
Given an integer A, representing number of vertices in a graph.
Also you are given a matrix of integers B of size N x 3 where N represents number of Edges in a Graph and Triplet (B[i][0], B[i][1], B[i][2]) implies there is an undirected edge between B[i][0] and B[i][1] and weight of that edge is B[i][2].
Find and return the weight of minimum weighted cycle and if there is no cycle return -1 instead.

NOTE: Graph may contain self loops but does not have duplicate edges.

Problem Constraints
1 <= A <= 1000
1 <= B[i][0], B[i][1] <= A
1 <= B[i][2] <= 100000

Input Format
The first argument given is the integer A.
The second argument given is the integer matrix B.

Output Format
Return the weight of minimum weighted cycle and if there is no cycle return -1 instead.

Example Input
Input 1:

	A = 4
	B = [  [1 ,2 ,2]
	       [2 ,3 ,3]
	       [3 ,4 ,1]
	       [4 ,1 ,4]
	       [1 ,3 ,15]  ]

Input 2:

	A = 3
	B = [  [1 ,2 ,2]
	       [2 ,3 ,3]  ]

Example Output
Output 1:
10

Output 2:
-1

Example Explanation
Explanation 1:

	Given graph forms 3 cycles
	1. 1 ---> 2 ---> 3 ---> 4 ---> 1 weight = 10
	2. 1 ---> 2 ---> 3 ---> 1 weight = 20
	3. 1 ---> 3---> 4 ---> 1 weight = 20
	so answer would be 10.

Explanation 2:

	Given graph forms 0 cycles so return -1.
*/
func GetMinWeightedCycle(A int, B [][]int) int {
	ans := math.MaxInt
	adjList := make([][]nodePair, A+1)
	//make adjList
	for i := range B {
		u := B[i][0]
		v := B[i][1]
		w := B[i][2]

		//if self loop, compare and save ans
		if u == v {
			ans = min(ans, w)
		}
		adjList[u] = append(adjList[u], nodePair{node: v, weight: w})
		adjList[v] = append(adjList[v], nodePair{node: u, weight: w})
	}

	//for each edge {u,v} and weight w, check if there is any alternate path exists with weight w'
	//if yes, save min of w+w' and ans
	for i := range B {
		u := B[i][0]
		v := B[i][1]
		w := B[i][2]

		if u == v {
			continue
		}
		cost := calculateCost(u, v, A+1, adjList)
		if cost != -1 {
			ans = min(ans, cost+w)
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}

// run Dijkshtra on src and destination node
func calculateCost(src, dest, N int, adjList [][]nodePair) int {
	dist := make([]int, N)
	for i := range dist {
		dist[i] = -1
	}
	mh := make(minHeap, 0)
	heap.Init(&mh)
	heap.Push(&mh, nodePair{node: src, weight: 0})

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
			if vert == src && v == dest {
				continue
			}
			//if not visited, push the minHeap
			if dist[v] == -1 {
				heap.Push(&mh, nodePair{node: v, weight: wgt + w})
			}
		}
	}
	return dist[dest]
}

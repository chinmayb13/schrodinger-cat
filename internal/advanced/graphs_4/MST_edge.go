package graphs4

import "sort"

/*
Problem Description

Given a undirected weighted graph with A nodes labelled from 1 to A with M edges given in a form of 2D-matrix B of size M * 3 where B[i][0] and B[i][1] denotes the two nodes connected by an edge of weight B[i][2].

For each edge check whether it belongs to any of the possible minimum spanning tree or not , return 1 if it belongs else return 0.

Return an one-dimensional binary array of size M denoting answer for each edge.

NOTE:

The graph may be disconnected in that case consider mst for each component.
No self-loops and no multiple edges present.
Answers in output array must be in order with the input array B output[i] must denote the answer of edge B[i][0] to B[i][1].

# Problem Constraints

1 <= A, M <= 3*105

1 <= B[i][0], B[i][1] <= A

1 <= B[i][1] <= 103

# Input Format

The first argument given is an integer A representing the number of nodes in the graph.

The second argument given is an matrix B of size M x 3 which represents the M edges such that there is a edge between node B[i][0] and node B[i][1] with weight B[i][2].

# Output Format

Return an one-dimensional binary array of size M denoting answer for each edge.

# Example Input

Input 1:

	A = 3
	B = [ [1, 2, 2]
	      [1, 3, 2]
	      [2, 3, 3]
	    ]

# Example Output

Output 1:

	[1, 1, 0]

# Example Explanation

Explanation 1:

	Edge (1, 2) with weight 2 is included in the MST           1
	                                                         /   \
	                                                        2     3
	Edge (1, 3) with weight 2 is included in the same MST mentioned above.
	Edge (2,3) with weight 3 cannot be included in any of the mst possible.
	So we will return [1, 1, 0]
*/
func FindMSTEdges(A int, B [][]int) []int {
	M := len(B)
	N := A
	ans := make([]int, M)
	container := make([][4]int, M)
	parent := make([]int, N+1)
	for i := 1; i <= N; i++ {
		parent[i] = i
	}

	for i := range B {
		for j := 0; j < 3; j++ {
			container[i][j] = B[i][j]
		}
		container[i][3] = i
	}

	//sort by weight
	sort.Slice(container, func(i, j int) bool {
		return container[i][2] < container[j][2]
	})

	i := 0
	for i < M {
		j := i
		//traverse edges of same weight
		for j < M && container[i][2] == container[j][2] {
			//if they're disjoint set, insert 1
			if find(container[j][0], parent) != find(container[j][1], parent) {
				ans[container[j][3]] = 1
			}
			j++
		}

		j = i
		//make union of the edge vertices
		for j < M && container[i][2] == container[j][2] {
			union(container[j][0], container[j][1], parent)
			j++
		}
		i = j
	}
	return ans

}

func find(v int, parent []int) int {
	if v == parent[v] {
		return v
	}
	parent[v] = find(parent[v], parent)
	return parent[v]
}

func union(v1, v2 int, parent []int) {
	p1 := find(v1, parent)
	p2 := find(v2, parent)
	if p1 == p2 {
		return
	}
	parent[p1] = p2
}

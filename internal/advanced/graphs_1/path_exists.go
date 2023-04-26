package graphs1

import "container/list"

/*
Problem Description
Given an directed graph having A nodes labelled from 1 to A containing M edges given by matrix B of size M x 2such that there is a edge directed from node

B[i][0] to node B[i][1].

Find whether a path exists from node 1 to node A.

Return 1 if path exists else return 0.

NOTE:

There are no self-loops in the graph.
There are no multiple edges between two nodes.
The graph may or may not be connected.
Nodes are numbered from 1 to A.
Your solution will run on multiple test cases. If you are using global variables make sure to clear them.

Problem Constraints
2 <= A <= 105

1 <= M <= min(200000,A*(A-1))

1 <= B[i][0], B[i][1] <= A

Input Format
The first argument given is an integer A representing the number of nodes in the graph.

The second argument given a matrix B of size M x 2 which represents the M edges such that there is a edge directed from node B[i][0] to node B[i][1].

Output Format
Return 1 if path exists between node 1 to node A else return 0.

Example Input
Input 1:

	A = 5
	B = [  [1, 2]
	       [4, 1]
	       [2, 4]
	       [3, 4]
	       [5, 2]
	       [1, 3] ]

Input 2:

	A = 5
	B = [  [1, 2]
	       [2, 3]
	       [3, 4]
	       [4, 5] ]

Example Output
Output 1:
0

Output 2:
1

Example Explanation
Explanation 1:
The given doens't contain any path from node 1 to node 5 so we will return 0.

Explanation 2:
Path from node1 to node 5 is ( 1 -> 2 -> 3 -> 4 -> 5 ) so we will return 1.
*/
func IsPathExist(A int, B [][]int) int {
	adjList := make([][]int, A+1)
	//build adjacency list
	for i := range B {
		adjList[B[i][0]] = append(adjList[B[i][0]], B[i][1])
	}
	hashSet := make(map[int]bool)
	q := newQueue()
	q.enqueueBack(1)
	hashSet[1] = true
	//do level order traversal to find the path
	for q.size() > 0 {
		elem := q.frontInt()
		q.dequeueFront()
		for i := range adjList[elem] {
			if adjList[elem][i] == A {
				return 1
			}
			if !hashSet[adjList[elem][i]] {
				hashSet[adjList[elem][i]] = true
				q.enqueueBack(adjList[elem][i])
			}
		}
	}
	return 0

}

type queue struct {
	*list.List
}

func newQueue() *queue {
	return &queue{
		List: list.New(),
	}
}

func (s *queue) enqueueBack(inp interface{}) {
	s.PushBack(inp)
}

func (s *queue) dequeueFront() {
	if s.Len() > 0 {
		s.Remove(s.Front())
	}
}

func (s *queue) front() interface{} {
	if s.Len() > 0 {
		return s.Front().Value

	}
	return nil
}

func (s *queue) size() int {
	return s.Len()
}

func (s *queue) frontInt() int {
	front := s.front()
	if front != nil {
		return front.(int)
	}
	return 0
}

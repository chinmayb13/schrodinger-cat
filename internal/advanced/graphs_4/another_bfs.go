package graphs4

import "container/list"

/*
Problem Description

Given a weighted undirected graph having A nodes, a source node C and destination node D.
Find the shortest distance from C to D and if it is impossible to reach node D from C then return -1.
You are expected to do it in Time Complexity of O(A + M).

Note:

There are no self-loops in the graph.
No multiple edges between two pair of vertices.
The graph may or may not be connected.
Nodes are Numbered from 0 to A-1.

Your solution will run on multiple testcases. If you are using global variables make sure to clear them.

Problem Constraints
1 <= A <= 105
0 <= B[i][0], B[i][1] < A
1 <= B[i][2] <= 2
0 <= C < A
0 <= D < A

Input Format

The first argument given is an integer A, representing the number of nodes.
The second argument given is the matrix B, where B[i][0] and B[i][1] are connected through an edge of weight B[i][2].
The third argument given is an integer C, representing the source node.
The fourth argument is an integer D, representing the destination node.
Note: B[i][2] will be either 1 or 2.

Output Format
Return the shortest distance from C to D. If it is impossible to reach node D from C then return -1.

Example Input

Input 1:
A = 6
B = [   [2, 5, 1]
        [1, 3, 1]
        [0, 5, 2]
        [0, 2, 2]
        [1, 4, 1]
        [0, 1, 1] ]
C = 3
D = 2

Input 2:
A = 2
B = [   [0, 1, 1]
    ]
C = 0
D = 1

Example Output
Output 1:
4

Output 2:
1

Example Explanation

Explanation 1:
The path to be followed will be:
    3 -> 1 (Edge weight : 1)
    1 -> 0 (Edge weight : 1)
    0 -> 2 (Edge weight : 2)
Total length of path = 1 + 1 + 2 = 4.

Explanation 2:
Path will be 0-> 1.
*/

// idea: to use temporary vertices whenever the edge weight is 2 and do BFS to find shortest distance
func GetShortestDistance(A int, B [][]int, C int, D int) int {
	adjList := make([][]int, max(A, len(B))<<1)
	visited := make([]bool, max(A, len(B))<<1)
	temp := A
	for i := range B {
		if B[i][2] == 2 {
			adjList[B[i][0]] = append(adjList[B[i][0]], temp)
			adjList[temp] = append(adjList[temp], B[i][0])

			adjList[B[i][1]] = append(adjList[B[i][1]], temp)
			adjList[temp] = append(adjList[temp], B[i][1])
			temp++
		} else {
			adjList[B[i][0]] = append(adjList[B[i][0]], B[i][1])
			adjList[B[i][1]] = append(adjList[B[i][1]], B[i][0])
		}
	}
	q := newQueue()
	q.enqueueBack(&costPair{node: C, cost: 0})
	visited[C] = true
	for q.size() > 0 {
		front := q.frontcostPair()
		q.dequeueFront()
		node := front.node
		cost := front.cost
		if node == D {
			return cost
		}
		for i := range adjList[node] {
			v := adjList[node][i]
			if !visited[v] {
				q.enqueueBack(&costPair{node: v, cost: cost + 1})
				visited[v] = true
			}
		}

	}
	return -1
}

type costPair struct {
	node int
	cost int
}

func max(inp1, inp2 int) int {
	if inp1 > inp2 {
		return inp1
	}
	return inp2
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

func (s *queue) frontcostPair() *costPair {
	front := s.front()
	if front != nil {
		return front.(*costPair)
	}
	return nil
}

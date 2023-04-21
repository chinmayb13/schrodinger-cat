package graphs3

import "container/list"

/*
Problem Description

Given a Tree of A nodes having A-1 edges. Each node is numbered from 1 to A where 1 is the root of the tree.
You are given Q queries. In each query, you will be given two integers L and X. Find the value of such node which lies at level L mod (MaxDepth + 1) and has value greater than or equal to X.
Answer to the query is the smallest possible value or -1, if all the values at the required level are smaller than X.

NOTE:

Level and Depth of the root is considered as 0.
It is guaranteed that each edge will be connecting exactly two different nodes of the tree.
Please read the input format for more clarification.

Problem Constraints

2 <= A, Q(size of array E and F) <= 105
1 <= B[i], C[i] <= A
1 <= D[i], E[i], F[i] <= 106

Input Format
The first argument is an integer A denoting the number of nodes.
The second and third arguments are the integer arrays B and C where for each i (0 <= i < A-1), B[i] and C[i] are the nodes connected by an edge.
The fourth argument is an integer array D, where D[i] denotes the value of the (i+1)th node
The fifth and sixth arguments are the integer arrays E and F where for each i (0 <= i < Q), E[i] denotes L and F[i] denotes X for ith query.

Output Format
Return an array of integers where the ith element denotes the answer to ith query.

Example Input
Input 1:
A = 5
B = [1, 4, 3, 1]
C = [5, 2, 4, 4]
D = [7, 38, 27, 37, 1]
E = [1, 1, 2]
F = [32, 18, 26]

Input 2:
A = 3
B = [1, 2]
C = [3, 1]
D = [7, 15, 27]
E = [1, 10, 1]
F = [29, 6, 26]

Example Output

Output 1:
[37, 37, 27]

Output 2:
[-1, 7, 27]

Example Explanation

Explanation 1:
      1[7]
     /    \
   5[1]  4[37]
        /    \
       2[38]  3[27]

Query 1:
L = 1, X = 32
Nodes for level 1 are 5, 4
Value of Node 5 = 1 < 32
Value of Node 4 = 37 >= 32
Ans = 37

Explanation 2:
      1[7]
     /    \
   2[15]  3[27]

Query 1:
L = 1, X = 6
Nodes for level 1 are 2, 3 having value 15 and 27 respectively.
Answer = -1 (Since no node is greater or equal to 29).

Query 1:
L = 10 % 2 = 0, X = 6
Nodes for level 0 is 1 having value 7.
Answer = 7.
*/

func GetMinLevelItem(A int, B []int, C []int, D []int, E []int, F []int) []int {
	adjList := make([][]int, A+1)
	//create adjacency list
	for i := range B {
		adjList[B[i]] = append(adjList[B[i]], C[i])
		adjList[C[i]] = append(adjList[C[i]], B[i])
	}

	//levels for retaining nodes in the same levels
	levels := make([][]int, A)

	//for marking the nodes as visited
	visited := make([]bool, A+1)

	//to identify max depth
	maxDepth := 0
	q := newQueue()
	q.enqueueBack(&pair{elem: 1, level: 0})
	levels[0] = append(levels[0], 1)
	visited[1] = true

	//do BFS
	for q.size() > 0 {
		front := q.frontPair()
		q.dequeueFront()
		elem := front.elem
		level := front.level
		if level > maxDepth {
			maxDepth = level
		}
		for i := range adjList[elem] {
			v := adjList[elem][i]
			//if not visited, enqueue, increase the level and mark the visited
			if !visited[v] {
				q.enqueueBack(&pair{elem: v, level: level + 1})
				levels[level+1] = append(levels[level+1], v)
				visited[v] = true
			}
		}
	}

	ans := make([]int, len(E))
	for i := range ans {
		ans[i] = -1
	}

	for i := range E {
		//target for comparison
		target := F[i]

		//get the level
		l := E[i] % (maxDepth + 1)

		//iterate through level and compare
		for j := range levels[l] {
			idx := levels[l][j] - 1
			if D[idx] >= target {
				if ans[i] == -1 || ans[i] > D[idx] {
					ans[i] = D[idx]
				}
			}
		}
	}
	return ans

}

type pair struct {
	elem  int
	level int
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

func (s *queue) frontPair() *pair {
	front := s.front()
	if front != nil {
		return front.(*pair)
	}
	return nil
}

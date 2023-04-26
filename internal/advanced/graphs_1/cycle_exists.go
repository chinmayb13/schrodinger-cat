package graphs1

/*
Problem Description
Given an directed graph having A nodes. A matrix B of size M x 2 is given which represents the M edges such that there is a edge directed from node B[i][0] to node B[i][1].

Find whether the graph contains a cycle or not, return 1 if cycle is present else return 0.

NOTE:
The cycle must contain atleast two nodes.
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
Return 1 if cycle is present else return 0.

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
1

Output 2:
0

Example Explanation
Explanation 1:
The given graph contain cycle 1 -> 3 -> 4 -> 1 or the cycle 1 -> 2 -> 4 -> 1

Explanation 2:
The given graph doesn't contain any cycle.
*/
func CycleExists(A int, B [][]int) int {
	adjList := make([][]int, A+1)
	//create adjacency list
	for i := range B {
		adjList[B[i][0]] = append(adjList[B[i][0]], B[i][1])
	}
	//create visited map
	vis := make(map[int]bool)
	for i := 1; i <= A; i++ {
		//loop on vertices and check if loop exists
		if !vis[i] && dfs(i, vis, adjList) == 1 {
			return 1
		}
	}
	return 0
}

func dfs(v int, vis map[int]bool, adjList [][]int) int {
	vis[v] = true
	for i := range adjList[v] {
		elem := adjList[v][i]
		if vis[elem] {
			return 1
		}
		dfs(elem, vis, adjList)
	}
	//we would want to delete the vertice key at the end of the func
	delete(vis, v)
	return 0
}

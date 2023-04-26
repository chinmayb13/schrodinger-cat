package graphs2

/*
Problem Description
You are given an undirected unweighted graph consisting of A vertices and M edges given in a form of 2D Matrix B of size M x 2 where (B[i][0], B][i][1]) denotes two nodes connected by an edge.
You have to write a number on each vertex of the graph. Each number should be 1, 2 or 3. The graph becomes Poisonous if for each edge the sum of numbers on vertices connected by this edge is odd.
Calculate the number of possible ways to write numbers 1, 2 or 3 on vertices so the graph becomes poisonous. Since this number may be large, return it modulo 998244353.

NOTE:
Note that you have to write exactly one number on each vertex.
The graph does not have any self-loops or multiple edges.
Nodes are labelled from 1 to A.

Problem Constraints
1 <= A <= 3*105
0 <= M <= 3*105
1 <= B[i][0], B[i][1] <= A
B[i][0] != B[i][1]

Input Format
First argument is an integer A denoting the number of nodes.
Second argument is an 2D Matrix B of size M x 2 denoting the M edges.

Output Format
Return one integer denoting the number of possible ways to write numbers 1, 2 or 3 on the vertices of given graph so it becomes Poisonous . Since answer may be large, print it modulo 998244353.

Example Input
Input 1:
A = 2
B = [  [1, 2]

	]

Input 2:
A = 4
B = [  [1, 2]

	    [1, 3]
	    [1, 4]
	    [2, 3]
	    [2, 4]
	    [3, 4]
	]

Example Output
Output 1:
4

Output 2:
0

Example Explanation
Explanation 1:
There are 4 ways to make the graph poisonous. i.e, writing number on node 1 and 2 as,

	[1, 2] , [3, 2], [2, 1] or [2, 3] repsectively.

Explanation 2:
There is no way to make the graph poisonous.
*/
func WaysToFormPoisonousGraph(A int, B [][]int) int {
	N := A
	adjList := make([][]int, N+1)
	//create adjacency list
	for i := range B {
		adjList[B[i][0]] = append(adjList[B[i][0]], B[i][1])
		adjList[B[i][1]] = append(adjList[B[i][1]], B[i][0])
	}

	//create color array
	colorArr := make([]int, N+1)
	for i := range colorArr {
		colorArr[i] = -1
	}
	ans := 1
	var mod int = 998244353
	for i := 1; i <= N; i++ {
		//for each connected component, check the number of zeros and ones
		//for each zero or one, place can be filled with either 1 or 3 while the other can be filled with 2
		//hence 2^zeros + 2^ones
		if colorArr[i] == -1 {
			p := &pair{zero: 1, one: 1}
			if !doDFS(i, 0, colorArr, adjList, p) {
				return 0
			}
			ans = (ans * (p.zero + p.one) % mod) % mod
		}
	}
	return ans
}

type pair struct {
	zero int
	one  int
}

func doDFS(vertice int, color int, colorArr []int, adjList [][]int, p *pair) bool {
	var mod int = 998244353
	colorArr[vertice] = color
	if color == 0 {
		p.zero = (p.zero << 1) % mod
	} else {
		p.one = (p.one << 1) % mod
	}
	for i := range adjList[vertice] {
		v := adjList[vertice][i]
		if colorArr[v] == -1 && !doDFS(v, color^1, colorArr, adjList, p) {
			return false
		} else if colorArr[v] != -1 && colorArr[v] == color {
			return false
		}
	}
	return true
}

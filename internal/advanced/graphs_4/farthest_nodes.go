package graphs4

/*
Find largest distance Given an arbitrary unweighted rooted tree which consists of N (2 <= N <= 40000) nodes.

The goal of the problem is to find largest distance between two nodes in a tree. Distance between two nodes is a number of edges on a path between the nodes (there will be a unique path between any pair of nodes since it is a tree).

The nodes will be numbered 0 through N - 1.

The tree is given as an array A, there is an edge between nodes A[i] and i (0 <= i < N). Exactly one of the i's will have A[i] equal to -1, it will be root node.

Problem Constraints
2 <= |A| <= 40000

Input Format
First and only argument is vector A

Output Format
Return the length of the longest path

Example Input
Input 1:
A = [-1, 0]

Input 2:
A = [-1, 0, 0]

Example Output
Output 1:
1

Output 2:
2

Example Explanation
Explanation 1:
Path is 0 -> 1.

Explanation 2:
Path is 1 -> 0 -> 2.
*/

func GetMaxDistance(A []int) int {
	N := len(A)
	edgeMap := make(map[int][]int)
	rootIdx := -1
	//build adjacancy list
	for i := range A {
		if A[i] == -1 {
			//save root node
			rootIdx = i
			continue
		}
		edgeMap[A[i]] = append(edgeMap[A[i]], i)
		edgeMap[i] = append(edgeMap[i], A[i])
	}
	//visited array
	visited := make([]bool, N)
	//to save farthest distance
	maxH := 0
	doDepthSearch(rootIdx, edgeMap, visited, &maxH)
	return maxH
}

func doDepthSearch(node int, edgeMap map[int][]int, visited []bool, maxH *int) (int, int) {
	visited[node] = true
	first := 0
	second := 0
	for i := range edgeMap[node] {
		n := edgeMap[node][i]
		//if not visited, do dfs
		if !visited[n] {
			h1, h2 := doDepthSearch(n, edgeMap, visited, maxH)
			//save the max two distances from current node
			height := max(h1, h2) + 1
			if height > first {
				second = first
				first = height
			} else if height > second {
				second = height
			}
		}
	}
	//compare if the farthest nodess have max distance than maxH
	if first+second > *maxH {
		*maxH = first + second
	}
	return first, second
}

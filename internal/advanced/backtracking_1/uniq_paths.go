package backtracking1

/*
Given a matrix of integers A of size N x M . There are 4 types of squares in it:

1. 1 represents the starting square.  There is exactly one starting square.
2. 2 represents the ending square.  There is exactly one ending square.
3. 0 represents empty squares we can walk over.
4. -1 represents obstacles that we cannot walk over.
Find and return the number of 4-directional walks from the starting square to the ending square, that walk over every non-obstacle square exactly once.

Note: Rows are numbered from top to bottom and columns are numbered from left to right.

Problem Constraints
2 <= N * M <= 20
-1 <= A[i] <= 2

Input Format
The first argument given is the integer matrix A.

Output Format
Return the number of 4-directional walks from the starting square to the ending square, that walk over every non-obstacle square exactly once.

Example Input
Input 1:
A = [   [1, 0, 0, 0]

	[0, 0, 0, 0]
	[0, 0, 2, -1]   ]

Input 2:
A = [   [0, 1]

	[2, 0]    ]

Example Output
Output 1:
2

Output 2:
0

Example Explanation
Explanation 1:
We have the following two paths:
1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)

Explanation 2:
Answer is evident here.
*/
func GetUniqPaths(A [][]int) int {
	var zeroCount int
	//hashMap to save the path being traced
	hashMap := make(map[int]bool)
	inp := A
	var begR, begC int
	N, M := len(inp), len(inp[0])
	size := N * M
	for i := range A {
		for j := range A[i] {
			switch A[i][j] {
			//count the number of zeros to be traversed in total
			case 0:
				zeroCount++
			case 1:
				//save the start indexes
				begR, begC = i, j
			}
		}
	}

	//start with star idx
	starIdx := begR*size + begC
	hashMap[starIdx] = true

	//count is -1 here, because we wouldn't want to include the start idx in the total hops, once we reach destination
	return traverse(inp, N, M, size, begR, begC, -1, zeroCount, hashMap)
}

func traverse(inp [][]int, N, M, size int, r, c, count, zeroCount int, hashMap map[int]bool) int {
	//invalid path
	if r < 0 || r == N || c < 0 || c == M || inp[r][c] == -1 {
		return 0
	}

	//if reached destination, check the hops are matching with total zero count
	if inp[r][c] == 2 {
		if count == zeroCount {
			return 1
		}
		return 0
	}

	//generate identifiers for next hops
	var up, right, down, left int
	upKey := (r-1)*size + c
	rightKey := r*size + c + 1
	downKey := (r+1)*size + c
	leftKey := r*size + c - 1

	//if the identifier is already present in the hashmap, it means the hop is already a part of the current path, no need to revisit

	if !hashMap[upKey] {
		hashMap[upKey] = true
		up = traverse(inp, N, M, size, r-1, c, count+1, zeroCount, hashMap)
		delete(hashMap, upKey)
	}
	if !hashMap[rightKey] {
		hashMap[rightKey] = true
		right = traverse(inp, N, M, size, r, c+1, count+1, zeroCount, hashMap)
		delete(hashMap, rightKey)
	}
	if !hashMap[downKey] {
		hashMap[downKey] = true
		down = traverse(inp, N, M, size, r+1, c, count+1, zeroCount, hashMap)
		delete(hashMap, downKey)
	}
	if !hashMap[leftKey] {
		hashMap[leftKey] = true
		left = traverse(inp, N, M, size, r, c-1, count+1, zeroCount, hashMap)
		delete(hashMap, leftKey)
	}

	//calculate sum of all traversal
	return up + right + down + left

}

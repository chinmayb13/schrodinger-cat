package graphs5

import "sort"

/*
Problem Description

Given a matrix C of integers, of dimension A x B.
For any given K, you are not allowed to travel between cells that have an absolute difference greater than K.
Return the minimum value of K such that it is possible to travel between any pair of cells in the grid through a path of adjacent cells.

NOTE:

Adjacent cells are those cells that share a side with the current cell.
# Problem Constraints
1 <= A, B <= 102
1 <= C[i][j] <= 109

# Input Format
The first argument given is A.
The second argument give is B.
The third argument given is the integer matrix C.

# Output Format
Return a single integer, the minimum value of K.

# Example Input
Input 1:

	A = 3
	B = 3
	C = [  [1, 5, 6]
	       [10, 7, 2]
	       [3, 6, 9]   ]

# Example Output

Output 1:
4

# Example Explanation

Explanation 1:
It is possible to travel between any pair of cells through a path of adjacent cells that do not have an absolute
difference in value greater than 4. e.g. : A path from (0, 0) to (2, 2) may look like this:
=> (0, 0) -> (0, 1) -> (1, 1) -> (2, 1) -> (2, 2)
*/

// idea: Use krustkal algo to find the max weight of any edge in a min spanning tree
func GetMaxPathWeight(A int, B int, C [][]int) int {
	ans := 0

	N := A * B

	//create parent array
	parent := make([]int, N)
	for i := range parent {
		parent[i] = i
	}

	//create items array
	var items [][]int
	for i := range C {
		for j := range C[i] {

			//save the row wise path difference
			if i+1 < A {
				items = append(items, []int{i*B + j, (i+1)*B + j, abs(C[i][j] - C[i+1][j])})
			}

			//save columnwise path difference
			if j+1 < B {
				items = append(items, []int{i*B + j, i*B + j + 1, abs(C[i][j] - C[i][j+1])})
			}
		}
	}

	//sort the items array based on the weight in ASC order
	sort.Slice(items, func(i, j int) bool {
		return items[i][2] < items[j][2]
	})

	edgeCount := 0
	for i := range items {
		pos1 := items[i][0]
		pos2 := items[i][1]

		//if disjoint set, increase the edge counter
		if union(pos1, pos2, parent) {
			edgeCount++

			//if edge counter has reached max, return the weight
			if edgeCount == N-1 {
				ans = items[i][2]
				break
			}
		}

	}

	return ans
}

func abs(inp int) int {
	if inp < 0 {
		return inp * -1
	}
	return inp
}

func getParent(inp int, parent []int) int {
	if inp == parent[inp] {
		return inp
	}
	parent[inp] = getParent(parent[inp], parent)
	return parent[inp]
}

func union(inp1, inp2 int, parent []int) bool {
	p1 := getParent(inp1, parent)
	p2 := getParent(inp2, parent)
	if p1 == p2 {
		return false
	}
	parent[p1] = p2
	return true
}

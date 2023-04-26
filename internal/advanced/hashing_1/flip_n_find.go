package hashing1

/*
Problem Description
Given a binary string A of size N. There are Q queries given by the array B of size Q*2.

Each query is given by :-
1 X :- Flip the bit of the X-th position in A
2 X :- Find the index of the nearest '1' from X. If there are multiple such indexes, return the one with the lower index. Return -1 if there are no '1's in A

Note :- We use 1-based indexing

Problem Constraints
1 <= N <= 105
1 <= Q <= 105
1 <= B[i][0] <= 2
1 <= B[i][1] <= N

Input Format
First argument A is a string.

Second argument B is a 2D array of integers describing the queries.

Output Format
Return an array of integers denoting the answers to each query of type 2.

Example Input
Input 1:
A = "10010"
B = [[1, 2]

	[2, 3]]

Input 2:
A = "010000100"
B = [[2, 5]

	[1, 7]
	[2, 9]]

Example Output
Output 1:
[2]

Output 2:
[7, 2]

Example Explanation
For Input 1:
After first query, A = "11010".
For second query, X = 3. Both index 2 and index 4 are at the same
distance but we choose the lower index.

For Input 2:
For first query, the index 2 is at a distance 3 and index 7 is at a distance 2. So we choose
index 7.
After second query, A = "010000000"
For third query, the only index with '1' is 2.
*/
func FlipAndFindNearest(A string, B [][]int) []int {
	var ans []int
	tm := &TreeMap{}
	for i := range A {
		if A[i] == '1' {
			tm.Put(i)
		}
	}
	for i := range B {
		key := B[i][1] - 1
		if B[i][0] == 1 {
			_, got := tm.Get(key)
			if got {
				tm.Delete(key)
			} else {
				tm.Put(key)
			}
		} else {
			lb, got1 := tm.LowerBound(key)
			if got1 && lb == key {
				ans = append(ans, key+1)
				continue
			}

			le, got2 := tm.LE(key)
			if got1 && got2 {
				if (lb - key) < (key - le) {
					ans = append(ans, lb+1)
				} else {
					ans = append(ans, le+1)
				}
			} else if got1 {
				ans = append(ans, lb+1)
			} else if got2 {
				ans = append(ans, le+1)
			} else {
				ans = append(ans, -1)
			}

		}
	}
	return ans
}

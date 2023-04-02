package dp5

/*
Problem Description
Given A, B, C find whether C is formed by the interleaving of A and B.

Problem Constraints
1 <= length(A), length(B) <= 100
1 <= length(C) <= 200

Input Format
The first argument of input contains a string, A.
The second argument of input contains a string, B.
The third argument of input contains a string, C.

Output Format
Return 1 if string C is formed by interleaving of A and B else 0.

Example Input
Input 1:
A = "aabcc"
B = "dbbca"
C = "aadbbcbcac"

Input 2:
A = "aabcc"
B = "dbbca"
C = "aadbbbaccc"

Example Output
Output 1:
1

Output 2:
0

Example Explanation
Explanation 1:
"aa" (from A) + "dbbc" (from B) + "bc" (from A) + "a" (from B) + "c" (from A)

Explanation 2:
It is not possible to get C by interleaving A and B.
*/
func IsInterleave(A string, B string, C string) int {
	if len(A)+len(B) != len(C) {
		return 0
	}
	src1 := A
	src2 := B
	target := C
	src1L := len(src1)
	src2L := len(src2)
	tarL := len(target)
	memo := make([][][]bool, src1L+1)
	for i := range memo {
		memo[i] = make([][]bool, src2L+1)
		for j := range memo[i] {
			memo[i][j] = make([]bool, tarL+1)
		}
	}

	//if src1 is exhausted, we would check if any prefix matches in src2 and target
	for j := 1; j <= src2L; j++ {
		if src2[j-1] != target[j-1] {
			break
		}
		memo[0][j][j] = true
	}

	//if src2 is exhausted, we would check if any prefix matches in src1 and target
	for i := 1; i <= src1L; i++ {
		if src1[i-1] != target[i-1] {
			break
		}
		memo[i][0][i] = true
	}

	for i := 1; i <= src1L; i++ {
		for j := 1; j <= src2L; j++ {
			for k := 1; k <= tarL; k++ {
				//if target character matches with src1 character
				// ans would depend on the prefix excluding the current src1 index
				//and current target index
				if target[k-1] == src1[i-1] {
					memo[i][j][k] = memo[i][j][k] || memo[i-1][j][k-1]
				}

				//if target character matches with src2 character
				// ans would depend on the prefix excluding the current src2 index
				//and current target index
				if target[k-1] == src2[j-1] {
					memo[i][j][k] = memo[i][j][k] || memo[i][j-1][k-1]
				}
			}
		}
	}

	if memo[src1L][src2L][tarL] {
		return 1
	}
	return 0
}

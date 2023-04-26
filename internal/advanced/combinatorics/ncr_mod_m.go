package combinatorics

import "github.com/chinmayb13/schrodinger-cat/internal/helper"

/*
Problem Description
Given three integers A, B, and C, where A represents n, B represents r, and C represents m, find and return the value of nCr % m where nCr % m = (n!/((n-r)!*r!))% m.
x! means factorial of x i.e. x! = 1 * 2 * 3... * x.

Problem Constraints
1 <= A * B <= 106

1 <= B <= A

1 <= C <= 106

Input Format
The first argument given is integer A ( = n).
The second argument given is integer B ( = r).
The third argument given is integer C ( = m).

Output Format
Return the value of nCr % m.

Example Input
Input 1:
A = 5
B = 2
C = 13

Input 2:
A = 6
B = 2
C = 13

Example Output
Output 1:
10

Output 2:
2

Example Explanation
Explanation 1:
The value of 5C2 % 11 is 10.

Explanation 2:
The value of 6C2 % 13 is 2.
*/

/*
approach: Create Pascal Triangle
-- -- -- -- -- 01 -- -- -- -- --
-- -- -- 01 -- 02 -- 01 -- -- --
-- -- 01 -- 03 -- 03 -- 01 -- --
-- 01 -- 04 -- 06 -- 04 -- 01 --
01 -- 05 -- 10 -- 10 -- 05 -- 01

Row 0 - 1
Row 1 - 1 1
Row 2 - 1 2 1
Row 3 - 1 3 3 1
Row 4 - 1 4 6 4 1
Row 5 - 1 5 10 10 5 1
*/
func GetnCrModM(A int, B int, C int) int {
	combination := make([][]int, A+1)
	for i := range combination {
		combination[i] = make([]int, B+1)
	}
	for i := 0; i <= A; i++ {
		for j := 0; j <= helper.Min(i, B); j++ {
			if j == 0 || j == i {
				combination[i][j] = 1
			} else {
				combination[i][j] = (combination[i-1][j-1] + combination[i-1][j]) % C
			}
		}
	}
	return combination[A][B]
}

func GetnCrModMAlt(A int, B int, C int) int {
	combination := make([]int, B+1)
	for i := 0; i <= A; i++ {
		if i <= B {
			combination[i] = 1
		}
		prev := make([]int, B+1)
		copy(prev, combination)
		for j := 1; j < i && j <= B; j++ {
			combination[j] = (combination[j] + prev[j-1]) % C
		}
	}
	return combination[B]
}

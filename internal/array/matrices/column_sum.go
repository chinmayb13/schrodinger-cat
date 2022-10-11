package matrices

/*
Problem Description
You are given a 2D integer matrix A, return a 1D integer array containing column-wise sums of original matrix.

Problem Constraints
1 <= A.size() <= 103
1 <= A[i].size() <= 103
1 <= A[i][j] <= 103

Input Format
First argument is a 2D array of integers.(2D matrix).

Output Format
Return an array conatining column-wise sums of original matrix.

Example Input
Input 1:
[1,2,3,4]
[5,6,7,8]
[9,2,3,4]


Example Output
Output 1:
{15,10,13,16}

Example Explanation
Explanation 1
Column 1 = 1+5+9 = 15
Column 2 = 2+6+2 = 10
Column 3 = 3+7+3 = 13
Column 4 = 4+8+4 = 16
*/
func GetColumnSum(A [][]int) []int {
	var columnSum []int
	rows := len(A)
	columns := len(A[0])
	for c := 0; c < columns; c++ {
		sum := 0
		for r := 0; r < rows; r++ {
			sum += A[r][c]
		}
		columnSum = append(columnSum, sum)
	}
	return columnSum
}

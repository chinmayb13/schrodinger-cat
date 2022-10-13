package slidingwindow

/*
Problem Description
Given an integer A, generate a square matrix filled with elements from 1 to A2 in spiral order and return the generated square matrix.

Problem Constraints
1 <= A <= 1000

Input Format
First and only argument is integer A

Output Format
Return a 2-D matrix which consists of the elements added in spiral order.

Example Input
Input 1:
1

Input 2:
2

Input 3:
5

Example Output
Output 1:
[ [1] ]

Output 2:
[ [1, 2], [4, 3] ]

Output 2:
[ [1, 2, 3, 4, 5], [16, 17, 18, 19, 6], [15, 24, 25, 20, 7], [14, 23, 22, 21, 8], [13, 12, 11, 10, 9] ]

Example Explanation
Explanation 1:
Only 1 is to be arranged.

Explanation 2:
1 --> 2
      |
      |
4<--- 3
*/
func GenerateSpiralMatrix(A int) [][]int {
	spiralMatrix := make([][]int, A)
	for i := range spiralMatrix {
		spiralMatrix[i] = make([]int, A)
	}

	//order of the square
	traverseLength := A
	row := 0
	col := 0
	value := 1
	//starting point for the perimeter
	boundaryStartRow := row
	boundaryStartCol := col
	//direction can be right down left and up, do it for all square with edge length > 1
	for direction := 0; traverseLength > 1; direction++ {
		//traverse from starting point till the second last cell in the edge
		for i := 1; i <= traverseLength-1; i++ {
			spiralMatrix[row][col] = value
			value++
			switch direction % 4 {
			case 0:
				col++
			case 1:
				row++
			case 2:
				col--
			case 3:
				row--
			}
		}
		//if reached starting point again, it means perimter has been traversed, move to inner square
		if row == boundaryStartRow && col == boundaryStartCol {
			traverseLength -= 2
			row++
			col++
			boundaryStartRow = row
			boundaryStartCol = col
		}

	}
	if traverseLength == 1 {
		spiralMatrix[row][col] = value
	}
	return spiralMatrix
}

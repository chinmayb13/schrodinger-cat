package hashing2

/*
Problem Description
Given two arrays of integers A and B describing a pair of (A[i], B[i]) coordinates in a 2D plane. A[i] describe x coordinates of the ith point in the 2D plane, whereas B[i] describes the y-coordinate of the ith point in the 2D plane.
Find and return the maximum number of points that lie on the same line.

Problem Constraints
1 <= (length of the array A = length of array B) <= 1000
-105 <= A[i], B[i] <= 105

Input Format
The first argument is an integer array A.
The second argument is an integer array B.

Output Format
Return the maximum number of points which lie on the same line.

Example Input
Input 1:
A = [-1, 0, 1, 2, 3, 3]
B = [1, 0, 1, 2, 3, 4]

Input 2:
A = [3, 1, 4, 5, 7, -9, -8, 6]
B = [4, -8, -3, -2, -1, 5, 7, -4]

Example Output
Output 1:
4

Output 2:
2

Example Explanation
Explanation 1:
The maximum number of point which lie on same line are 4, those points are {0, 0}, {1, 1}, {2, 2}, {3, 3}.

Explanation 2:
Any 2 points lie on a same line.
*/
func GetMaxPointsInLine(A []int, B []int) int {
	var ans, gcd, currAns int
	var centrePoints, verticalPoints, numerator, denominator int
	var lineMap map[pair]int

	length := len(A)
	for i := 0; i < length; i++ {
		centrePoints = 0
		verticalPoints = 0
		currAns = 0
		lineMap = make(map[pair]int)

		for j := i + 1; j < length; j++ {
			//num and denom for calculating slope from ith point
			numerator = B[j] - B[i]
			denominator = A[j] - A[i]
			//overlapping points
			if numerator == 0 && denominator == 0 {
				centrePoints++
				//points forming a line || y-axis
			} else if denominator == 0 {
				verticalPoints++
				//getGCD to save the slope in integer format
			} else {
				gcd = getGCD(numerator, denominator)
				numerator /= gcd
				denominator /= gcd
				lineMap[pair{numerator, denominator}]++
				currAns = max(currAns, lineMap[pair{numerator, denominator}])
			}
		}
		currAns = max(currAns, verticalPoints)
		ans = max(ans, currAns+centrePoints+1)

	}
	return ans
}

func getGCD(inp1, inp2 int) int {
	if inp2 == 0 {
		return inp1
	}
	return getGCD(inp2, inp1%inp2)
}

func max(inp1, inp2 int) int {
	if inp1 > inp2 {
		return inp1
	}
	return inp2
}

type pair struct {
	first, second int
}

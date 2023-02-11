package stack2

/*
Problem Description
Given an array of integers A.
A represents a histogram i.e A[i] denotes the height of the ith histogram's bar. Width of each bar is 1.
Find the area of the largest rectangle formed by the histogram.

Problem Constraints
1 <= |A| <= 100000
1 <= A[i] <= 1000000000

Input Format
The only argument given is the integer array A.

Output Format
Return the area of the largest rectangle in the histogram.

Example Input
Input 1:
A = [2, 1, 5, 6, 2, 3]

Input 2:
A = [2]

Example Output
Output 1:
10

Output 2:
2

Example Explanation
Explanation 1:
The largest rectangle has area = 10 unit. Formed by A[3] to A[4].

Explanation 2:
Largest rectangle has area 2.
*/
func LargestRectangleArea(A []int) int {
	ans := 0
	nsl := getNearestSmallerFromLeft(A)
	nsr := getNearestSmallerFromRight(A)
	for i := 0; i < len(A); i++ {
		ans = max(ans, A[i]*(nsr[i]-nsl[i]-1))
	}
	return ans
}

func getNearestSmallerFromLeft(inp []int) []int {
	nsl := make([]int, len(inp))
	candidateList := newStack()
	for i := range inp {
		for candidateList.size() > 0 && inp[candidateList.topInt()] >= inp[i] {
			candidateList.pop()
		}
		if candidateList.size() > 0 {
			nsl[i] = candidateList.topInt()
		} else {
			nsl[i] = -1
		}
		candidateList.push(i)
	}
	return nsl
}

func getNearestSmallerFromRight(inp []int) []int {
	nsr := make([]int, len(inp))
	candidateList := newStack()
	for i := len(inp) - 1; i >= 0; i-- {
		for candidateList.size() > 0 && inp[candidateList.topInt()] >= inp[i] {
			candidateList.pop()
		}
		if candidateList.size() > 0 {
			nsr[i] = candidateList.topInt()
		} else {
			nsr[i] = len(inp)
		}
		candidateList.push(i)
	}
	return nsr
}

func max(inp1, inp2 int) int {
	if inp1 > inp2 {
		return inp1
	}
	return inp2
}

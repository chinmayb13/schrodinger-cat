package twopointers

/*
Problem Description
Given a binary array A, find the maximum sequence of continuous 1's that can be formed by replacing at-most B zeroes.
For this problem, return the indices of maximum continuous series of 1s in order.
If there are multiple possible solutions, return the sequence which has the minimum start index.

Problem Constraints
0 <= B <= 105
1 <= size(A) <= 105
0 <= A[i] <= 1

Input Format
First argument is an binary array A.
Second argument is an integer B.

Output Format
Return an array of integers denoting the indices(0-based) of 1's in the maximum continuous series.

Example Input
Input 1:
A = [1 1 0 1 1 0 0 1 1 1 ]
B = 1

Input 2:
A = [1, 0, 0, 0, 1, 0, 1]
B = 2

Example Output
Output 1:
[0, 1, 2, 3, 4]

Output 2:
[3, 4, 5, 6]

Example Explanation
Explanation 1:
Flipping 0 present at index 2 gives us the longest continous series of 1's i.e subarray [0:4].

Explanation 2:
Flipping 0 present at index 3 and index 5 gives us the longest continous series of 1's i.e subarray [3:6].
*/
func GetMaxOnesCount(A []int, B int) []int {
	var length, maxLength int
	var ans []int
	var beg int
	zeroCount := 0
	left, right := 0, 0
	for i := range A {
		//increase count when 0 encountered
		if A[i] == 0 {
			zeroCount++
		}

		//if zeroes are more than the limit, move left ptr to right till zeroes get in limit
		for zeroCount > B {
			if A[left] == 0 {
				zeroCount--
			}
			left++
		}

		//if length is more than max so far, save the ans
		length = right - left + 1
		if length > maxLength {
			maxLength = length
			beg = left
		}
		right++

	}

	if maxLength > 0 {
		ans = make([]int, maxLength)
		for i := 0; i < maxLength; i++ {
			ans[i] = beg
			beg++
		}
	}
	return ans

}

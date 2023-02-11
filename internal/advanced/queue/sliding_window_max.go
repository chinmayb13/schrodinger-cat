package queue

/*
Problem Description
Given an array of integers A. There is a sliding window of size B, moving from the very left of the array to the very right. You can only see the B numbers in the window. Each time the sliding window moves rightwards by one position. You have to find the maximum for each window.
Return an array C, where C[i] is the maximum value in the array from A[i] to A[i+B-1].
Refer to the given example for clarity.
NOTE: If B > length of the array, return 1 element with the max of the array.

Problem Constraints
1 <= |A|, B <= 106

Input Format
The first argument given is the integer array A.
The second argument given is the integer B.

Output Format
Return an array C, where C[i] is the maximum value of from A[i] to A[i+B-1].

Example Input
Input 1:
A = [1, 3, -1, -3, 5, 3, 6, 7]
B = 3

Input 2:
A = [1, 2, 3, 4, 2, 7, 1, 3, 6]
B = 6

Example Output
Output 1:
[3, 3, 5, 5, 6, 7]

Output 2:
[7, 7, 7, 7]

Example Explanation
Explanation 1:
Window position     | Max
--------------------|-------
[1 3 -1] -3 5 3 6 7 | 3
1 [3 -1 -3] 5 3 6 7 | 3
1 3 [-1 -3 5] 3 6 7 | 5
1 3 -1 [-3 5 3] 6 7 | 5
1 3 -1 -3 [5 3 6] 7 | 6
1 3 -1 -3 5 [3 6 7] | 7

Explanation 2:
Window position     | Max
--------------------|-------
[1 2 3 4 2 7] 1 3 6 | 7
1 [2 3 4 2 7 1] 3 6 | 7
1 2 [3 4 2 7 1 3] 6 | 7
1 2 3 [4 2 7 1 3 6] | 7
*/
func GetSlidingMaximum(A []int, B int) []int {
	var ans []int
	size := len(A)
	candidateList := newQueue()
	for i := 0; i < min(size, B); i++ {
		//always keep the candidate list in decreasing order
		//such that the front always points to the maximum
		for candidateList.size() > 0 && A[candidateList.backInt()] <= A[i] {
			candidateList.dequeueBack()
		}
		candidateList.enqueueBack(i)
	}
	//fetch the maximum
	ans = append(ans, A[candidateList.frontInt()])
	l, r := 1, B
	for r < size {
		//if the maximum is out of the window, remove from candidate list
		if l-1 == candidateList.frontInt() {
			candidateList.dequeueFront()
		}
		//repeat the same behavior for incoming elements
		for candidateList.size() > 0 && A[candidateList.backInt()] <= A[r] {
			candidateList.dequeueBack()
		}
		candidateList.enqueueBack(r)
		ans = append(ans, A[candidateList.frontInt()])
		l++
		r++
	}
	return ans

}

func min(inp1, inp2 int) int {
	if inp1 < inp2 {
		return inp1
	}
	return inp2
}

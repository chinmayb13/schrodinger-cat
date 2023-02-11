package stack2

/*
Problem Description
Given an integer array A of size N. You have to generate it's all subarrays having a size greater than 1.
Then for each subarray, find Bitwise XOR of its maximum and second maximum element.
Find and return the maximum value of XOR among all subarrays.

Problem Constraints
2 <= N <= 105
1 <= A[i] <= 107

Input Format
The only argument is an integer array A.

Output Format
Return an integer, i.e., the maximum value of XOR of maximum and 2nd maximum element among all subarrays.

Example Input
Input 1:
A = [2, 3, 1, 4]

Input 2:
A = [1, 3]

Example Output
Output 1:
7

Outnput 2:
2

Example Explanation
Explanation 1:
All subarrays of A having size greater than 1 are:
Subarray            XOR of maximum and 2nd maximum no.
1. [2, 3]           1
2. [2, 3, 1]        1
3. [2, 3, 1, 4]     7
4. [3, 1]           2
5. [3, 1, 4]        7
6. [1, 4]           5
So maximum value of Xor among all subarrays is 7.

Explanation 2:
Only subarray is [1, 3] and XOR of maximum and 2nd maximum is 2.
*/
func GetMaxXor(A []int) int {
	return max(getLeftMaxXor(A), getrightMaxXor(A))
}

//we are taking NGL because, because we want to fix the current element as the max
//and try xoring with smaller elements before popping

//idea: the xor can be maximised by taking the max and comparing it with smaller elements
func getLeftMaxXor(inp []int) int {
	leftMaxXor := 0
	candidateList := newStack()
	for i := range inp {
		for candidateList.size() > 0 && inp[candidateList.topInt()] <= inp[i] {
			leftMaxXor = max(leftMaxXor, inp[i]^inp[candidateList.topInt()])
			candidateList.pop()
		}
		candidateList.push(i)
	}
	return leftMaxXor
}

//get nearest smaller element from right
func getrightMaxXor(inp []int) int {
	rightMaxXor := 0
	candidateList := newStack()
	for i := len(inp) - 1; i >= 0; i-- {
		for candidateList.size() > 0 && inp[candidateList.topInt()] <= inp[i] {
			rightMaxXor = max(rightMaxXor, inp[i]^inp[candidateList.topInt()])
			candidateList.pop()
		}
		candidateList.push(i)
	}
	return rightMaxXor
}

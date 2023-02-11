package stack2

/*
Problem Description
Given an array of integers A.
value of a array = max(array) - min(array).
Calculate and return the sum of values of all possible subarrays of A modulo 109+7.

Problem Constraints
1 <= |A| <= 100000
1 <= A[i] <= 1000000

Input Format
The first and only argument given is the integer array A.

Output Format
Return the sum of values of all possible subarrays of A modulo 109+7.

Example Input
Input 1:
A = [1]

Input 2:
A = [4, 7, 3, 8]

Example Output
Output 1:
0

Output 2:
26

Example Explanation
Explanation 1:
Only 1 subarray exists. Its value is 0.

Explanation 2:
value ( [4] ) = 4 - 4 = 0
value ( [7] ) = 7 - 7 = 0
value ( [3] ) = 3 - 3 = 0
value ( [8] ) = 8 - 8 = 0
value ( [4, 7] ) = 7 - 4 = 3
value ( [7, 3] ) = 7 - 3 = 4
value ( [3, 8] ) = 8 - 3 = 5
value ( [4, 7, 3] ) = 7 - 3 = 4
value ( [7, 3, 8] ) = 8 - 3 = 5
value ( [4, 7, 3, 8] ) = 8 - 3 = 5
sum of values % 10^9+7 = 26
*/
func GetSumMaxMinusMin(A []int) int {
	var ans int
	var mod int = 1e9 + 7
	nsl := getNSL(A)
	nsr := getNSR(A)
	ngl := getNGL(A)
	ngr := getNGR(A)
	for i := range A {
		//max returns contribution of A[i] where A[i] is maximum
		max := (((i - ngl[i]) * (ngr[i] - i)) % mod * A[i]) % mod
		//min returns contribution of A[i] where A[i] is minimum
		min := (((i - nsl[i]) * (nsr[i] - i)) % mod * A[i]) % mod
		ans = (ans + (max-min+mod)%mod) % mod
	}
	return ans
}

//get nearest smaller element from left
func getNSL(inp []int) []int {
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

//get nearest smaller element from right
func getNSR(inp []int) []int {
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

//get nearest greater element from left
func getNGL(inp []int) []int {
	ngl := make([]int, len(inp))
	candidateList := newStack()
	for i := range inp {
		for candidateList.size() > 0 && inp[candidateList.topInt()] <= inp[i] {
			candidateList.pop()
		}
		if candidateList.size() > 0 {
			ngl[i] = candidateList.topInt()
		} else {
			ngl[i] = -1
		}
		candidateList.push(i)
	}
	return ngl
}

//get nearest greater element from right
func getNGR(inp []int) []int {
	ngr := make([]int, len(inp))
	candidateList := newStack()
	for i := len(inp) - 1; i >= 0; i-- {
		for candidateList.size() > 0 && inp[candidateList.topInt()] <= inp[i] {
			candidateList.pop()
		}
		if candidateList.size() > 0 {
			ngr[i] = candidateList.topInt()
		} else {
			ngr[i] = len(inp)
		}
		candidateList.push(i)
	}
	return ngr
}

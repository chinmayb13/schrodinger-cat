package hashing1

/*
Misha likes finding all Subarrays of an Array. Now she gives you an array A of N elements and told you to find the number of subarrays of A, that have unique elements.
Since the number of subarrays could be large, return value % 109 +7.

# Problem Constraints

1 <= N <= 105
1 <= A[i] <= 106

Input Format
The only argument given is an Array A, having N integers.

Output Format
Return the number of subarrays of A, that have unique elements.

Example Input
Input 1:
A = [1, 1, 3]

Input 2:
A = [2, 1, 2]

Example Output
Output 1:
4

Output 1:
5

Example Explanation
Explanation 1:
Subarrays of A that have unique elements only:
[1], [1], [1, 3], [3]

Explanation 2:
Subarrays of A that have unique elements only:
[2], [1], [2, 1], [1, 2], [2]
*/
func GetSubArrayCount(A []int) int {
	var mod int = 1e9 + 7
	left, right := 0, 0
	saCount := 0
	idxMap := make(map[int]int)
	for i := range A {
		idx, ok := idxMap[A[i]]
		//if duplicate is found, move the left ptr to remove the previously visited duplicate
		//only when the idx is greater than the left
		if ok && left <= idx {
			left = idx + 1
		}
		saCount = (saCount + right - left + 1) % mod
		idxMap[A[i]] = i
		right++
	}
	return saCount
}

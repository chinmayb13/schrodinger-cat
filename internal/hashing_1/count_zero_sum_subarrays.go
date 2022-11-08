package hashing1

/*
Problem Description
Given an array A of N integers.

Find the count of the subarrays in the array which sums to zero. Since the answer can be very large, return the remainder on dividing the result with 109+7


Problem Constraints
1 <= N <= 105
-109 <= A[i] <= 109


Input Format
Single argument which is an integer array A.


Output Format
Return an integer.


Example Input
Input 1:
A = [1, -1, -2, 2]

Input 2:
A = [-1, 2, -1]

Example Output
Output 1:
3

Output 2:
1

Example Explanation
Explanation 1:
The subarrays with zero sum are [1, -1], [-2, 2] and [1, -1, -2, 2].

Explanation 2:
The subarray with zero sum is [-1, 2, -1].
*/
//approach: subarray with [L,R] sum will be 0 if the sum at the R index is same as sum at L-1th index
func GetZeroSumSubarraysCount(A []int) int {
	countMap := map[int64]int{0: 1}
	var sum int64 = 0
	for i := range A {
		sum += int64(A[i])
		countMap[sum]++
	}
	var count int64 = 0
	for _, v := range countMap {
		if v > 1 {
			count += int64(((v - 1) * v / 2) % (1e9 + 7))
		}
	}
	return int(count % (1e9 + 7))
}

func GetZeroSumSubarraysCountAlt(A []int) int {
	var divisor int = 1e9 + 7
	countMap := make(map[int64]int)
	var count int64 = 0
	var sum int64 = 0
	for i := range A {
		sum += int64(A[i])
		if sum == 0 {
			count++
		}
		freq, ok := countMap[sum]
		if !ok {
			countMap[sum] = 1
		} else {
			count = (count + (int64(freq))%int64(divisor)) % int64(divisor)
			countMap[sum]++
		}
	}
	return int(count % (1e9 + 7))
}

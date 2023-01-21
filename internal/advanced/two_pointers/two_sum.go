package twopointers

/*
Given a sorted array of integers (not necessarily distinct) A and an integer B, find and return how many pair of integers ( A[i], A[j] ) such that i != j have sum equal to B.
Since the number of such pairs can be very large, return number of such pairs modulo (109 + 7).

Problem Constraints
1 <= |A| <= 100000
1 <= A[i] <= 10^9
1 <= B <= 10^9

Input Format
The first argument given is the integer array A.
The second argument given is integer B.

Output Format
Return the number of pairs for which sum is equal to B modulo (10^9+7).

Example Input
Input 1:
A = [1, 1, 1]
B = 2

Input 2:
A = [1, 1]
B = 2

Example Output
Output 1:
3

Output 2:
1

Example Explanation
Explanation 1:
Any two pairs sum up to 2.

Explanation 2:
only pair (1, 2) sums up to 2.
*/
func GetTwoSumPairs(A []int, B int) int {
	var mod int = 1e9 + 7
	pairCount := 0
	l, r := 0, len(A)-1
	for l < r {
		sum := A[l] + A[r]
		if sum > B {
			r--
		} else if sum < B {
			l++
		} else {

			lCount := 0
			lVal := A[l]
			for l < r && A[l] == lVal {
				lCount++
				l++
			}

			rCount := 0
			rVal := A[r]
			for l <= r && A[r] == rVal {
				rCount++
				r--
			}

			if lVal == rVal {
				addendum := ((lCount + rCount) * (lCount + rCount - 1)) >> 1
				pairCount = (pairCount + (addendum)%mod) % mod
			} else {
				pairCount = (pairCount + (lCount*rCount)%mod) % mod
			}
		}
	}
	return pairCount
}

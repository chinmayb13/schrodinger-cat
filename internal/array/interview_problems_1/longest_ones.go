package interviewproblems1

/*
Given a binary string A. It is allowed to do at most one swap between any 0 and 1. Find and return the length of the longest consecutive 1’s that can be achieved.

Input Format
The only argument given is string A.

Output Format
Return the length of the longest consecutive 1’s that can be achieved.

Constraints
1 <= length of string <= 1000000
A contains only characters 0 and 1.

For Example
Input 1:
    A = "111000"
Output 1:
    3

Input 2:
    A = "111011101"
Output 2:
    7
*/
func GetLongestConsecutiveOnes(A string) int {
	countOne := 0
	//count the number of ones in the input array
	for i := range A {
		if A[i] == '1' {
			countOne++
		}
	}

	//if the input doesn't contain any zero
	if countOne == len(A) {
		return countOne
	}

	longestOneCount := 0
	for i := range A {
		//encounter a zero
		if A[i] == '0' {
			leftCountOne := 0
			//count the number of 1s in the left hand side
			for leftIdx := i - 1; leftIdx >= 0; leftIdx-- {
				if A[leftIdx] == '0' {
					break
				}
				leftCountOne++
			}

			rightCountOne := 0
			//count the number of 1s in the right hand side
			for rightIdx := i + 1; rightIdx < len(A); rightIdx++ {
				if A[rightIdx] == '0' {
					break
				}
				rightCountOne++
			}

			eitherCountOne := leftCountOne + rightCountOne

			//if there exists a 1 outside the subarray
			if countOne > eitherCountOne {
				totalLocalCount := eitherCountOne + 1
				//check if the subarray is longest
				if totalLocalCount > longestOneCount {
					longestOneCount = totalLocalCount
				}
			} else {
				//check if the subarray is longest
				if eitherCountOne > longestOneCount {
					longestOneCount = eitherCountOne
				}
			}
		}
	}
	return longestOneCount
}

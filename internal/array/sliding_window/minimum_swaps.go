package slidingwindow

/*
Problem Description
Given an array of integers A and an integer B, find and return the minimum number of swaps required to bring all the numbers less than or equal to B together.
Note: It is possible to swap any two elements, not necessarily consecutive.

Problem Constraints
1 <= length of the array <= 100000
-109 <= A[i], B <= 109

Input Format
The first argument given is the integer array A.
The second argument given is the integer B.

Output Format
Return the minimum number of swaps.

# Example Input

Input 1:
A = [1, 12, 10, 3, 14, 10, 5]
B = 8

	Input 2:

A = [5, 17, 100, 11]
B = 20

Example Output
Output 1:
2

Output 2:
1

# Example Explanation

Explanation 1:
A = [1, 12, 10, 3, 14, 10, 5]
After swapping  12 and 3, A => [1, 3, 10, 12, 14, 10, 5].
After swapping  the first occurence of 10 and 5, A => [1, 3, 5, 12, 14, 10, 10].
Now, all elements less than or equal to 8 are together.

Explanation 2:
A = [5, 17, 100, 11]
After swapping 100 and 11, A => [5, 17, 11, 100].
Now, all elements less than or equal to 20 are together.
*/
func GetMinimumSwaps(A []int, B int) int {

	subArraySize := 0
	//Calculate the number of elements less than equal to B to calculate subArray size
	for i := range A {
		if A[i] <= B {
			subArraySize++
		}
	}

	swapCount := 0
	//Calculate swaps needed in the 0th subArray to accumulate all elements less than equal to B
	for i := 0; i < subArraySize; i++ {
		if A[i] > B {
			swapCount++
		}
	}
	minSwapCount := swapCount

	//Calculate  swaps needed in the subsequent subArrays and find the minimum swap
	for l, r := 1, subArraySize; r < len(A); l, r = l+1, r+1 {
		if A[l-1] <= B {
			swapCount++
		}
		if A[r] <= B {
			swapCount--
		}

		if swapCount < minSwapCount {
			minSwapCount = swapCount
		}
	}
	return minSwapCount
}

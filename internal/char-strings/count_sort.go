package charstrings

/*
Problem Description
Given an array A. Sort this array using Count Sort Algorithm and return the sorted array.

Problem Constraints
1 <= |A| <= 105
1 <= A[i] <= 105

Input Format
The first argument is an integer array A.

Output Format
Return an integer array that is the sorted array A.

Example Input
Input 1:
A = [1, 3, 1]

Input 2:
A = [4, 2, 1, 3]

Example Output
Output 1:
[1, 1, 3]

Output 2:
[1, 2, 3, 4]

Example Explanation
For Input 1:
The array in sorted order is [1, 1, 3].

For Input 2:
The array in sorted order is [1, 2, 3, 4].
*/
func CountSort(A []int) []int {
	countArr := make([]int, 1e5)
	//make array to count the presence of each element in the range
	for i := range A {
		//since range is from 1 t0 1e5, we are mapping 1 to 0th index and so on
		countArr[A[i]-1]++
	}
	k := 0
	for i := range countArr {
		//for each element, run loop as many times as the count
		for j := 0; j < countArr[i]; j++ {
			//assign the element back to original array
			A[k] = i + 1
			k++
		}
	}
	return A
}

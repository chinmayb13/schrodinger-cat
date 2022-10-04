package carryforward

/*
Problem Description
Given an array A, find the size of the smallest subarray such that it contains at least one occurrence of the maximum value of the array
and at least one occurrence of the minimum value of the array.

Problem Constraints
1 <= |A| <= 2000



Input Format
First and only argument is vector A



Output Format
Return the length of the smallest subarray which has at least one occurrence of minimum and maximum element of the array



Example Input
Input 1:
A = [1, 3]

Input 2:
A = [2]


Example Output

Output 1:
2

Output 2:
1


Example Explanation

Explanation 1:
Only choice is to take both elements.

Explanation 2:
Take the whole array.
*/

func GetSubArrLength(A []int) int {
	arrLength := len(A)
	if len(A) == 1 {
		return 1
	}

	//since array is of atleast 1 length
	max := A[0]
	min := A[0]
	for i := 1; i < arrLength; i++ {
		if A[i] > max {
			max = A[i]
		}
		if A[i] < min {
			min = A[i]
		}
	}

	//if all the elements are same, then return 1 as the element would be min and max itself
	if min == max {
		return 1
	}

	//start with non present index
	lastSeenMaxIdx, lastSeenMinIdx := -1, -1

	//we know that atleast the whole array would consist of min and max element
	subArrLength := arrLength

	for i := range A {
		if A[i] == max {
			if lastSeenMinIdx != -1 { //min element has been found atleast once
				length := i - lastSeenMinIdx + 1
				if length < subArrLength {
					subArrLength = length
				}
			}
			lastSeenMaxIdx = i //update the new found max index
		} else if A[i] == min {
			if lastSeenMaxIdx != -1 { //max element has been found atleast once
				length := i - lastSeenMaxIdx + 1
				if length < subArrLength {
					subArrLength = length
				}
			}
			lastSeenMinIdx = i //update the new found min index
		}
	}
	return subArrLength
}

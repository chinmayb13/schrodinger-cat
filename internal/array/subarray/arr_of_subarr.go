package subarray

/*
Problem Description
You are given an array A of N integers.
Return a 2D array consisting of all the subarrays of the array


Problem Constraints
1 <= N <= 100
1 <= A[i] <= 105


Input Format
First argument A is an array of integers.


Output Format
Return a 2D array of integers


Example Input
Input 1:
A = [1, 2, 3]
Input 2:
A = [5, 2, 1, 4]


Example Output
Output 1:
[[1], [1, 2], [1, 2, 3], [2], [2, 3], [3]]

Output 2:
[[5], [5, 2], [5, 2, 1], [5, 2, 1, 4], [2], [2, 1], [2, 1, 4], [1], [1, 4], [4]]


Example Explanation
For Input 1:
All the subarrays of the array are returned. There are a total of 6 subarrays.
For Input 2:
All the subarrays of the array are returned. There are a total of 10 subarrays.
*/

func Get2DSubArray(A []int) [][]int {
	var arrSubArray [][]int
	arrLength := len(A)
	for i := 0; i < arrLength; i++ {
		var sumArr []int
		for j := i; j < arrLength; j++ {
			sumArr = append(sumArr, A[j])
			arrSubArray = append(arrSubArray, sumArr)
		}
	}
	return arrSubArray
}

// func Get2DSubArray(A []int) [][]int {
// 	arrLength := len(A)
// 	subArrLength := arrLength * (arrLength + 1) / 2
// 	arrSubArray := make([][]int, subArrLength)
// 	mdArrCount := 0
// 	for l := range A {
// 		for r := l; r < len(A); r++ {
// 			var subArray []int
// 			for i := l; i <= r; i++ {
// 				subArray = append(subArray, A[i])
// 			}
// 			arrSubArray[mdArrCount]=subArray
// 			mdArrCount++
// 		}
// 	}
// 	return arrSubArray
// }

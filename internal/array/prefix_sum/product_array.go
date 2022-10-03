package prefixsum

/*
Given an array of integers A, find and return the product array of the same size where the ith element of the product array will be equal to the product of all the elements divided by the ith element of the array.

Note: It is always possible to form the product array with integer (32 bit) values. Solve it without using the division operator.

Input Format
The only argument given is the integer array A.

Output Format
Return the product array.

Constraints
2 <= length of the array <= 1000
1 <= A[i] <= 10

For Example

Input 1:
    A = [1, 2, 3, 4, 5]
Output 1:
    [120, 60, 40, 30, 24]

Input 2:
    A = [5, 1, 10, 1]
Output 2:
    [10, 50, 5, 50]
*/

func GetProductArray(A []int) []int {
	arrLength := len(A)
	var productArr []int
	prefixSum := getPrefixSum(A)
	suffixSum := getSuffixSum(A)
	for i := range A {
		//nothing on the left
		if i == 0 {
			productArr = append(productArr, suffixSum[i+1])
			//nothing on the right
		} else if i == arrLength-1 {
			productArr = append(productArr, prefixSum[i-1])
		} else {
			productArr = append(productArr, prefixSum[i-1]*suffixSum[i+1])
		}
	}
	return productArr

}

func getPrefixSum(A []int) []int {
	var prefixSum []int
	for i := range A {
		if i == 0 {
			prefixSum = append(prefixSum, A[i])
		} else {
			prefixSum = append(prefixSum, prefixSum[i-1]*A[i])
		}
	}
	return prefixSum
}

func getSuffixSum(A []int) []int {
	arrLength := len(A)
	suffixSum := make([]int, arrLength)
	for i := arrLength - 1; i >= 0; i-- {
		if i == arrLength-1 {
			suffixSum[i] = A[i]
		} else {
			suffixSum[i] = suffixSum[i+1] * A[i]
		}
	}
	return suffixSum
}

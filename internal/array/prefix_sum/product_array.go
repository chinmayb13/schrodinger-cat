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

// func GetProductArray(A []int) []int {
// 	arrLength := len(A)
// 	var productArr []int
// 	prefixProduct := getPrefixProduct(A)
// 	suffixProduct := getSuffixProduct(A)
// 	for i := range A {
// 		//nothing on the left
// 		if i == 0 {
// 			productArr = append(productArr, suffixProduct[i+1])
// 			//nothing on the right
// 		} else if i == arrLength-1 {
// 			productArr = append(productArr, prefixProduct[i-1])
// 		} else {
// 			productArr = append(productArr, prefixProduct[i-1]*suffixProduct[i+1])
// 		}
// 	}
// 	return productArr

// }

func GetProductArray(A []int) []int {
	arrLength := len(A)
	var productArr []int
	prefixProduct := getPrevPrefixProduct(A)
	suffixProduct := getPostSuffixProduct(A)
	for i := range A {
		//nothing on the left
		if i == 0 {
			productArr = append(productArr, suffixProduct[i])
			//nothing on the right
		} else if i == arrLength-1 {
			productArr = append(productArr, prefixProduct[i])
		} else {
			productArr = append(productArr, prefixProduct[i]*suffixProduct[i])
		}
	}
	return productArr
}

// func getPrefixProduct(A []int) []int {
// 	var prefixProduct []int
// 	for i := range A {
// 		if i == 0 {
// 			prefixProduct = append(prefixProduct, A[i])
// 		} else {
// 			prefixProduct = append(prefixProduct, prefixProduct[i-1]*A[i])
// 		}
// 	}
// 	return prefixProduct
// }

// func getSuffixProduct(A []int) []int {
// 	arrLength := len(A)
// 	suffixProduct := make([]int, arrLength)
// 	for i := arrLength - 1; i >= 0; i-- {
// 		if i == arrLength-1 {
// 			suffixProduct[i] = A[i]
// 		} else {
// 			suffixProduct[i] = suffixProduct[i+1] * A[i]
// 		}
// 	}
// 	return suffixProduct
// }

func getPrevPrefixProduct(A []int) []int {
	prefixProduct := make([]int, len(A))
	prefixProduct[0] = 1
	for i := 1; i < len(A); i++ {
		prefixProduct[i] = prefixProduct[i-1] * A[i-1]
	}
	return prefixProduct
}

func getPostSuffixProduct(A []int) []int {
	suffixProduct := make([]int, len(A))
	suffixProduct[len(A)-1] = 1
	for i := len(A) - 2; i >= 0; i-- {
		suffixProduct[i] = suffixProduct[i+1] * A[i+1]
	}
	return suffixProduct
}

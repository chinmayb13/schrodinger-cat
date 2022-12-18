package primenumbers

import "math"

/*
Problem Description
Groot has an array A of size N. Boring, right? Groot thought so too, so he decided to construct another array B of the same size and defined elements of B as:

B[i] = factorial of A[i] for every i such that 1<= i <= N.

factorial of a number X denotes (1 x 2 x 3 x 4......(X-1) x (X)).
Now Groot is interested in the total number of non-empty subsequences of array B such that every element in the subsequence has the same non-empty set of prime factors.

Since the number can be huge, return it modulo 109 + 7.

NOTE: A set is a data structure having only distinct elements. E.g : the set of prime factors of Y=12 will be {2,3} and not {2,2,3}

Problem Constraints
1 <= N <= 105
1 <= A[i] <= 106
Your code will run against a maximum of 5 test cases.

Input Format
Only argument is an integer array A of size N.

Output Format
Return an integer deonting the total number of non-empty subsequences of array B such that every element in the subsequence has the same set of prime factors modulo 109+7.

Example Input
Input 1:
A = [2, 3, 2, 3]

Input 2:
A = [2, 3, 4]

Example Output
Output 1:
6

Output 2:
4

Example Explanation
Explanation 1:
Array B will be : [2, 6, 2, 6]
The total possible subsequences are 6 : [2], [2], [2, 2], [6], [6], [6, 6].

Input 2:
Array B will be : [2, 6, 24]
The total possible subsequences are 4 : [2], [6], [24], [6, 24].
*/
func GetPrimeSubsequenceCount(A []int) int {
	ans := 0
	var mod int = 1e9 + 7
	hashMap := make(map[int]int)
	max := getMax(A)
	sieveArray := buildSieveArrayBinary(max + 1)
	enrichSieveArrayBinary(sieveArray, max)
	for i := 0; i < len(A); i++ {
		if A[i] > 1 {
			hashMap[sieveArray[A[i]]]++
		}
	}
	for _, v := range hashMap {
		ans = (ans + ((1<<v)-1)%mod) % mod
	}
	return ans
}

func buildSieveArrayBinary(size int) []int {
	sieveArray := make([]int, size)
	for i := 2; i < size; i++ {
		sieveArray[i] = 1
	}
	return sieveArray
}

func enrichSieveArrayBinary(sieveArray []int, inp int) {
	sqrt := int(math.Sqrt(float64(inp)))
	for i := 2; i <= sqrt; i++ {
		if sieveArray[i] > 0 {
			for j := i * i; j <= inp; j += i {
				sieveArray[j] = 0
			}
		}
	}

	for i := 1; i < len(sieveArray); i++ {
		sieveArray[i] += sieveArray[i-1]
	}

}

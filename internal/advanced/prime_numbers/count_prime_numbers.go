package primenumbers

import "math"

/*
Problem Description
Given an integer A. Find the list of all prime numbers in the range [1, A].

Problem Constraints
1 <= A <= 106

Input Format
Only argument A is an integer.

Output Format
Return a sorted array of prime number in the range [1, A].

Example Input
Input 1:
A = 7

Input 2:
A = 12

Example Output
Output 1:
[2, 3, 5, 7]

Output 2:
[2, 3, 5, 7, 11]

Example Explanation
For Input 1:
The prime numbers from 1 to 7 are 2, 3, 5 and 7.

For Input 2:
The prime numbers from 1 to 12 are 2, 3, 5, 7 and 11.
*/
func GetPrimeCount(A int) []int {
	var primeNumber []int
	sieveArray := buildSieveArrayBool(A + 1)
	enrichSieveArrayBool(sieveArray, A)
	for i := range sieveArray {
		if sieveArray[i] {
			primeNumber = append(primeNumber, i)
		}
	}
	return primeNumber
}

func buildSieveArrayBool(size int) []bool {
	sieveArray := make([]bool, size)
	for i := 2; i < size; i++ {
		sieveArray[i] = true
	}
	return sieveArray
}

func enrichSieveArrayBool(sieveArray []bool, inp int) {
	sqrt := int(math.Sqrt(float64(inp)))
	for i := 2; i <= sqrt; i++ {
		if sieveArray[i] {
			for j := i * i; j <= inp; j += i {
				sieveArray[j] = false
			}
		}
	}
}

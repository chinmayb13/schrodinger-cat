package bitmanipulate1

/*
Problem Description
Given an array of integers, every element appears thrice except for one, which occurs once.
Find that element that does not appear thrice.
NOTE: Your algorithm should have a linear runtime complexity.
Could you implement it without using extra memory?

Problem Constraints
2 <= A <= 5*106
0 <= A <= INTMAX

Input Format
First and only argument of input contains an integer array A.

Output Format
Return a single integer.

Example Input
Input 1:
A = [1, 2, 4, 3, 3, 2, 2, 3, 1, 1]

Input 2:
A = [0, 0, 0, 1]

Example Output
Output 1:
4

Output 2:
1

Example Explanation
Explanation 1:
4 occurs exactly once in Input 1.
1 occurs exactly once in Input 2.
*/
func GetSingleNumber(A []int) int {
	soloElement := 0
	max := -1
	for i := range A {
		if A[i] > max {
			max = A[i]
		}
	}
	bitCount := geBitCount(max)
	for j := 0; j < bitCount; j++ {
		setBits := 0
		for i := range A {
			if (A[i] & (1 << j)) > 0 {
				setBits++
			}
		}
		if (setBits % 3) > 0 {
			soloElement |= (1 << j)
		}
	}
	return soloElement

}

func geBitCount(A int) int {
	numOfBits := 0
	for A > 0 {
		numOfBits++
		A = A >> 1
	}
	return numOfBits
}

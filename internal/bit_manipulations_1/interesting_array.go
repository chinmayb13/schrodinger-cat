package bitmanipulations1

/*
Problem Description
You have an array A with N elements. We have two types of operation available on this array :
We can split an element B into two elements, C and D, such that B = C + D.
We can merge two elements, P and Q, to one element, R, such that R = P ^ Q i.e., XOR of P and Q.
You have to determine whether it is possible to convert array A to size 1, containing a single element equal to 0 after several splits and/or merge?

Problem Constraints
1 <= N <= 100000
1 <= A[i] <= 106

Input Format
The first argument is an integer array A of size N.

Output Format
Return "Yes" if it is possible otherwise return "No".

Example Input
Input 1:
A = [9, 17]

Input 2:
A = [1]

Example Output
Output 1:
Yes

Output 2:
No

Example Explanation

Explanation 1:
Following is one possible sequence of operations -
1) Merge i.e 9 XOR 17 = 24
2) Split 24 into two parts each of size 12
3) Merge i.e 12 XOR 12 = 0
As there is only 1 element i.e 0. So it is possible.

Explanation 2:
There is no possible way to make it 0.
*/

/*
SOLUTION
If any element in array is even then it can be made 0. Split that element in two equal parts of  Ai/2 and Ai/2. XOR of two equal numbers is zero. Therefore this strategy makes an element 0.

If  any element is odd. Split it in two part: 1, Ai-1. Since Ai-1 is even, it can be made 0 by above strategy. Therefore an odd element can reduce its size to 1.

Two odd elements can therefore be made 0 by following above strategy and finally XOR them (i.e. 1) as 1 XOR 1 = 0.

Therefore if number of odd elements in array is even, then answer is possible. Otherwise an element of value 1 will be left and it is not possible to satisfy the condition.
*/
func IsAnInterestingArray(A []int) string {
	oddCount := 0
	for i := range A {
		if (A[i] & 1) == 1 {
			oddCount++
		}
	}
	if (oddCount & 1) == 1 {
		return "No"
	}
	return "Yes"
}

func IsAnInterestingArrayAlt(A []int) string {
	xor := 0
	for i := range A {
		xor ^= A[i]
	}
	if (xor & 1) == 1 {
		return "No"
	}
	return "Yes"
}

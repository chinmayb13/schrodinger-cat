package stack2

/*
Problem Description
Given a stack of integers A, sort it using another stack.
Return the array of integers after sorting the stack using another stack.

Problem Constraints
1 <= |A| <= 5000
0 <= A[i] <= 109

Input Format
The only argument is a stack given as an integer array A.

Output Format
Return the array of integers after sorting the stack using another stack.

Example Input
Input 1:
A = [5, 4, 3, 2, 1]

Input 2:
A = [5, 17, 100, 11]

Example Output
Output 1:
[1, 2, 3, 4, 5]

Output 2:
[5, 11, 17, 100]

Example Explanation
Explanation 1:
Just sort the given numbers.

Explanation 2:
Just sort the given numbers.
*/
func SortStack(A []int) []int {
	primStack := newStack()
	secStack := newStack()
	for i := range A {
		//idea: top of primary stack should always be smallest
		//if top of primary stack is smaller than A[i], dump the data to secondary stack
		for primStack.size() > 0 && primStack.topInt() < A[i] {
			secStack.push(primStack.topInt())
			primStack.pop()
		}
		//push the element into primary stack
		primStack.push(A[i])
		//dump the data back to primary stack from secondary stack
		for secStack.size() > 0 {
			primStack.push(secStack.topInt())
			secStack.pop()
		}
	}
	for i := 0; i < len(A); i++ {
		A[i] = primStack.topInt()
		primStack.pop()
	}
	return A
}

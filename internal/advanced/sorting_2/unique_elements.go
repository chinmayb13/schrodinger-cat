package sorting2

import "sort"

/*
Problem Description
You are given an array A of N elements. You have to make all elements unique. To do so, in one step you can increase any number by one.
Find the minimum number of steps.

Problem Constraints
1 <= N <= 105
1 <= A[i] <= 109

Input Format
The only argument given is an Array A, having N integers.

Output Format
Return the minimum number of steps required to make all elements unique.

Example Input
Input 1:
A = [1, 1, 3]

Input 2:
A = [2, 4, 5]

Example Output
Output 1:
1

Output 2:
0

Example Explanation
Explanation 1:
We can increase the value of 1st element by 1 in 1 step and will get all unique elements. i.e [2, 1, 3].

Explanation 2:
All elements are distinct.
*/

func GetMinStepsToUniquify(A []int) int {
	var minSteps int
	sort.Ints(A)
	for i := 1; i < len(A); i++ {
		//if the prev element is greater or equal to current element
		if A[i] <= A[i-1] {
			//steps to make the current element greater than the previous
			steps := A[i-1] - A[i] + 1
			minSteps += steps
			A[i] += steps
		}
	}
	return minSteps
}

func GetMinStepsToUniquifyAlt(A []int) int {
	var minSteps int
	sort.Ints(A)
	for i := 1; i < len(A); i++ {
		if A[i] == A[i-1] {
			steps := stepsToUniquify(A, A[i], i+1)
			minSteps += steps
			A = insertionSort(A, A[i]+steps, i+1)
		}
	}
	return minSteps
}

func insertionSort(inp []int, val, beg int) []int {
	ptr := len(inp) - 1
	inp = append(inp, val)
	if val < inp[ptr] {
		for (ptr >= beg) && (inp[ptr] > val) {
			inp[ptr+1] = inp[ptr]
			ptr--
		}
		inp[ptr+1] = val
	}
	return inp
}

func stepsToUniquify(inp []int, val, beg int) int {
	for i := beg; i < len(inp); i++ {
		if inp[i]-inp[i-1] > 1 {
			return inp[i-1] + 1 - val
		}
	}
	return inp[len(inp)-1] + 1 - val
}

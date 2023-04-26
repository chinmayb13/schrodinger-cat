package queue

/*
Problem Description
Given an array A of both positive and negative integers.
Your task is to compute the sum of minimum and maximum elements of all sub-array of size B.
NOTE: Since the answer can be very large, you are required to return the sum modulo 109 + 7.

Problem Constraints
1 <= size of array A <= 105
-109 <= A[i] <= 109
1 <= B <= size of array

Input Format
The first argument denotes the integer array A.
The second argument denotes the value B

Output Format
Return an integer that denotes the required value.

Example Input
Input 1:
A = [2, 5, -1, 7, -3, -1, -2]
B = 4

Input 2:
A = [2, -1, 3]
B = 2

Example Output
Output 1:
18

Output 2:
3

Example Explanation
Explanation 1:
Subarrays of size 4 are :

	[2, 5, -1, 7],   min + max = -1 + 7 = 6
	[5, -1, 7, -3],  min + max = -3 + 7 = 4
	[-1, 7, -3, -1], min + max = -3 + 7 = 4
	[7, -3, -1, -2], min + max = -3 + 7 = 4
	Sum of all min & max = 6 + 4 + 4 + 4 = 18

Explanation 2:
Subarrays of size 2 are :

	[2, -1],   min + max = -1 + 2 = 1
	[-1, 3],   min + max = -1 + 3 = 2
	Sum of all min & max = 1 + 2 = 3
*/
func GetMinMaxSum(A []int, B int) int {
	var mod int = 1e9 + 7
	max := getMaxElementSum(A, B)
	min := getMinElementSum(A, B)
	return (max + min + mod) % mod

}

// getMaxElementSum calculates sum of all maxs in all subarrays of window size
func getMaxElementSum(inp []int, window int) int {
	var mod int = 1e9 + 7
	var ans int
	size := len(inp)
	queue := newQueue()
	for i := 0; i < window; i++ {
		for queue.size() > 0 && inp[queue.backInt()] <= inp[i] {
			queue.dequeueBack()
		}
		queue.enqueueBack(i)
	}
	ans = (ans + inp[queue.frontInt()]) % mod
	l, r := 1, window
	for r < size {
		if l-1 == queue.frontInt() {
			queue.dequeueFront()
		}

		for queue.size() > 0 && inp[queue.backInt()] <= inp[r] {
			queue.dequeueBack()
		}
		queue.enqueueBack(r)
		ans = (ans + inp[queue.frontInt()]) % mod
		l++
		r++
	}
	return ans
}

// getMinElementSum calculates sum of all mins in all subarrays of window size
func getMinElementSum(inp []int, window int) int {
	var mod int = 1e9 + 7
	var ans int
	size := len(inp)
	queue := newQueue()
	for i := 0; i < window; i++ {
		for queue.size() > 0 && inp[queue.backInt()] >= inp[i] {
			queue.dequeueBack()
		}
		queue.enqueueBack(i)
	}
	ans = (ans + inp[queue.frontInt()]) % mod
	l, r := 1, window
	for r < size {
		if l-1 == queue.frontInt() {
			queue.dequeueFront()
		}

		for queue.size() > 0 && inp[queue.backInt()] >= inp[r] {
			queue.dequeueBack()
		}
		queue.enqueueBack(r)
		ans = (ans + inp[queue.frontInt()]) % mod
		l++
		r++
	}
	return ans
}

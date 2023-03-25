package dp2

import "math"

/*
Problem Description
Max Heap is a special kind of complete binary tree in which, for every node, the value present in that node is greater than the value present in its children nodes.

Find the number of distinct Max Heap that can be made from A distinct integers.

In short, you have to ensure the following properties for the max heap :

Heap has to be a complete binary tree ( A complete binary tree is a binary tree in which every level, except possibly the last, is completely filled, and all nodes are as far left as possible.)
Every node is greater than all its children.

NOTE: If you want to know more about Heaps, please visit this link. Return your answer modulo 109 + 7.

Problem Constraints
1 <= A <= 100

Input Format
The first and only argument is an integer A.

Output Format
Return an integer denoting the number of distinct Max Heap.

Example Input
Input 1:
A = 4

Input 2:
A = 10

Example Output
Output 1:
3

Output 2:
3360

Example Explanation
Explanation 1:
Let us take 1, 2, 3, 4 as our 4 distinct integers
Following are the 3 possible max heaps from these 4 numbers :
     4           4                     4
   /  \         / \                   / \
  3    2   ,   2   3      and        3   1
 /            /                     /
1            1                     2

Explanation 2:
Number of distinct heaps possible with 10 distinct integers = 3360.
*/

func GetHeapWays(A int) int {
	inp := int64(A)
	memo := make([]int64, inp+1)
	for i := range memo {
		memo[i] = -1
	}
	return int(maxHeapWays(inp, memo))
}

func maxHeapWays(inp int64, memo []int64) int64 {
	if inp <= 2 {
		return 1
	}

	if memo[inp] != -1 {
		return memo[inp]
	}

	var mod int64 = 1e9 + 7

	//get the count of left max heap nodes
	leftHeapElements := getLeftHeapElements(inp)

	//get the count of right max heap nodes
	rightHeapElements := inp - 1 - leftHeapElements

	//find ways of selecting left max heap nodes
	//(n-1)C(l)
	ans := getCombination(inp-1, leftHeapElements)

	//once combinations are found, recursively call it for left and right max heap
	ans = (ans * maxHeapWays(leftHeapElements, memo)) % mod
	ans = (ans * maxHeapWays(rightHeapElements, memo)) % mod
	memo[inp] = ans
	return ans
}

func getLeftHeapElements(inp int64) int64 {
	height := int(math.Floor(math.Log2(float64(inp)) + 1))
	leafNodeCount := inp - (1<<(height-1) - 1)
	if leafNodeCount >= (1 << (height - 2)) {
		return 1<<(height-1) - 1
	}
	return 1<<(height-2) - 1 + leafNodeCount
}

func inverse(inp int64) int64 {
	var mod int64 = 1e9 + 7
	pow := mod - 2
	var ans int64 = 1
	for pow > 0 {
		if pow&1 > 0 {
			ans = (ans * inp) % mod
			pow--
		} else {
			inp = (inp * inp) % mod
			pow = pow >> 1
		}
	}
	return ans
}

func fact(inp int64) int64 {
	var mod int64 = 1e9 + 7
	var ans, i int64 = 1, 1
	for i <= inp {
		ans = (ans * i) % mod
		i++
	}
	return ans
}

func getCombination(n, r int64) int64 {
	var mod int64 = 1e9 + 7
	return (((fact(n) * inverse(fact(n-r))) % mod) * inverse(fact(r))) % mod
}

package tries2

/*
Problem Description
Given an array of integers A, find and return the maximum result of A[i] XOR A[j], where i, j are the indexes of the array.

Problem Constraints
1 <= length of the array <= 100000
0 <= A[i] <= 109

Input Format
The only argument given is the integer array A.

Output Format
Return an integer denoting the maximum result of A[i] XOR A[j].

Example Input
Input 1:
A = [1, 2, 3, 4, 5]

Input 2:
A = [5, 17, 100, 11]

Example Output
Output 1:
7

Output 2:
117

Example Explanation
Explanation 1:
Maximum XOR occurs between element of indicies(0-based) 1 and 4 i.e. 2 ^ 5 = 7.

Explanation 2:
Maximum XOR occurs between element of indicies(0-based) 1 and 2 i.e. 17 ^ 100 = 117.
*/
func GetMaxXor(A []int) int {
	//get the number of bits of the max number
	bits := getBits(A)
	if bits == 0 {
		return 0
	}
	//create trie of the bits
	root := getTrie(A, bits)
	ans := 0
	for i := range A {
		cur := root
		localAns := 0
		for k := bits - 1; k >= 0; k-- {
			bit := (1 << k) & A[i]
			//if kth bit is set, look if the trie has kth bit as 0
			if bit > 0 {
				if cur.left != nil {
					localAns += 1 << k
					cur = cur.left
				} else {
					cur = cur.right
				}
				//if kth bit is unset, look if the trie has kth bit as 1
			} else {
				if cur.right != nil {
					localAns += 1 << k
					cur = cur.right
				} else {
					cur = cur.left
				}
			}
		}
		ans = max(ans, localAns)
	}
	return ans
}

func getTrie(A []int, bits int) *btNode {
	root := newbtNode()
	for i := range A {
		cur := root
		for k := bits - 1; k >= 0; k-- {
			if (A[i] & (1 << k)) > 0 {
				if cur.right == nil {
					cur.right = newbtNode()
				}
				cur = cur.right
			} else {
				if cur.left == nil {
					cur.left = newbtNode()
				}
				cur = cur.left
			}
		}

	}
	return root
}

func max(args ...int) int {
	ans := 0
	for i := range args {
		if args[i] > ans {
			ans = args[i]
		}
	}
	return ans
}

package tries2

/*
Problem Description
Given an array of integers A of size N.
A triplet (i, j, k), i < j <= k is called a power triplet if A[i] ^ A[i+1] ^....A[j-1] = A[j] ^.....^A[k].
Where, ^ denotes bitwise xor.
Return the count of all possible power triplets. Since the answer could be large return answer % 109 +7.

Problem Constraints
1 <= N <= 100000
1 <= A[i] <= 100000

Input Format
The first argument given is the integer array A.

Output Format
Return the count of all possible power triplets % 109 + 7.

Example Input
Input 1:
A = [5, 2, 7]

Input 2:
A = [1, 2, 3]

Example Output
Output 1:
2

Output 2:
2

Example Explanation
Explanation 1:
All possible power triplets are:
    1. (1, 2, 3) ->  A[1] = A[2] ^ A[3]
    2. (1, 3, 3) ->  A[1] ^ A[2] = A[3]

Explanation 2:
All possible power triplets are:
    1. (1, 2, 3) ->  A[1] = A[2] ^ A[3]
    2. (1, 3, 3) ->  A[1] ^ A[2] = A[3]
*/
func GetTripletCount(A []int) int {
	xor := 0
	ans := 0
	bits := getBits(A)
	root := newbtNode()
	//insert 0 at -1 position because the first subarray's xor could be zero
	root.insertNode(bits, 0, -1)
	for i := range A {
		//xor the elements iteratively
		xor ^= A[i]
		//find how many existing indexes have the given xor value, for each idx, count the triplets and add
		ans = (ans + root.search(bits, xor, i)) % (1e9 + 7)
		root.insertNode(bits, xor, i)
	}
	return ans
}

func (root *btNode) search(bits, val, idx int) int {
	cur := root
	for k := bits - 1; k >= 0; k-- {
		if val&(1<<k) > 0 {
			if cur.right == nil {
				return 0
			}
			cur = cur.right
		} else {
			if cur.left == nil {
				return 0
			}
			cur = cur.left
		}
	}
	sum := 0
	for i := range cur.arr {
		sum = (sum + idx - cur.arr[i] - 1) % (1e9 + 7)
	}
	return sum
}

func (root *btNode) insertNode(bits, val, idx int) {
	cur := root
	for k := bits - 1; k >= 0; k-- {
		if (val & (1 << k)) > 0 {
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
	cur.arr = append(cur.arr, idx)
}

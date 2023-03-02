package tries2

/*
Problem Description
Given an array, A of integers of size N. Find the subarray AL, AL+1, AL+2, ... AR with 1<=L<=R<=N, which has maximum XOR value.
NOTE: If there are multiple subarrays with the same maximum value, return the subarray with minimum length. If the length is the same, return the subarray with the minimum starting index.

Problem Constraints
1 <= N <= 100000
0 <= A[i] <= 109

Input Format
First and only argument is an integer array A.

Output Format
Return an integer array B of size 2. B[0] is the starting index(1-based) of the subarray and B[1] is the ending index(1-based) of the subarray.

Example Input
Input 1:
A = [1, 4, 3]

Input 2:
A = [8]

Example Output
Output 1:
[2, 3]

Output 2:
[1, 1]

Example Explanation
Explanation 1:
There are 6 possible subarrays of A:
subarray            XOR value
[1]                     1
[4]                     4
[3]                     3
[1, 4]                  5 (1^4)
[4, 3]                  7 (4^3)
[1, 4, 3]               6 (1^4^3)
[4, 3] subarray has maximum XOR value. So, return [2, 3].

Explanation 2:
There is only one element in the array. So, the maximum XOR value is equal to 8 and the only possible subarray is [1, 1].
*/
func GetMaxXorSubarray(A []int) []int {
	bits := getBits(A)
	if bits == 0 {
		return []int{1, 1}
	}
	root := newbtNode()
	xor, maxXor := 0, 0
	//initialize beg and end with array first and last index
	beg, end := 0, len(A)-1
	//insert 0 at the beginning, since the max xor could be the first subarray itself
	root.insert(bits, 0, -1)
	for i := 0; i < len(A); i++ {
		xor ^= A[i]
		//get the max xor value by xoring the overall xor with any other previously calculated xor
		inf := root.getMaxXor(bits, xor)
		//if the xor is max or the subarray length is smaller
		if inf.xor > maxXor || inf.xor == maxXor && (i-inf.beg) < (end-beg) {
			maxXor = inf.xor
			beg = inf.beg + 1
			end = i
		}
		//insert the overall xor
		root.insert(bits, xor, i)
	}
	return []int{beg + 1, end + 1}
}

func (root *btNode) insert(bits, val, idx int) {
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
	cur.value = idx
}

func (root *btNode) getMaxXor(bits, val int) info {
	cur := root
	xor := 0
	for k := bits - 1; k >= 0; k-- {
		if (val & (1 << k)) > 0 {
			if cur.left != nil {
				xor += 1 << k
				cur = cur.left
			} else {
				cur = cur.right
			}
		} else {
			if cur.right != nil {
				xor += 1 << k
				cur = cur.right
			} else {
				cur = cur.left
			}
		}
	}
	return info{
		xor: xor,
		beg: cur.value,
	}
}

package heaps1

import "reflect"

/*
Problem Description
We have a list A of points (x, y) on the plane. Find the B closest points to the origin (0, 0).
Here, the distance between two points on a plane is the Euclidean distance.
You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in.)

NOTE: Euclidean distance between two points P1(x1, y1) and P2(x2, y2) is sqrt( (x1-x2)2 + (y1-y2)2 ).

Problem Constraints
1 <= B <= length of the list A <= 105
-105 <= A[i][0] <= 105
-105 <= A[i][1] <= 105

Input Format
The argument given is list A and an integer B.

Output Format
Return the B closest points to the origin (0, 0) in any order.

Example Input
Input 1:
 A = [
       [1, 3],
       [-2, 2]
     ]
 B = 1

Input 2:
 A = [
       [1, -1],
       [2, -1]
     ]
 B = 1

Example Output
Output 1:
[ [-2, 2] ]

Output 2:
[ [1, -1] ]

Example Explanation
Explanation 1:
The Euclidean distance will be sqrt(10) for point [1,3] and sqrt(8) for point [-2,2].
So one closest point will be [-2,2].

Explanation 2:
The Euclidean distance will be sqrt(2) for point [1,-1] and sqrt(5) for point [2,-1].
So one closest point will be [1,-1].
*/
func GetBClosesPoints(A [][]int, B int) [][]int {
	var ans [][]int
	//create min heap
	createMinPointHeap(A)
	for i := 0; i < B; i++ {
		//get min element
		ans = append(ans, A[0])
		//remove min element by swapping with last element and removing the last element
		A[0] = A[len(A)-1]
		A = A[:len(A)-1]
		//heapify again
		minPointHeapify(0, A)
	}
	return ans
}

func getLfc(idx int, inp [][]int) []int {
	lc := (idx << 1) + 1
	if lc >= len(inp) {
		return nil
	}
	return inp[lc]
}

func getRtc(idx int, inp [][]int) []int {
	rc := (idx << 1) + 2
	if rc >= len(inp) {
		return nil
	}
	return inp[rc]
}

func minPointHeapify(idx int, inp [][]int) {
	if len(inp) < 2 {
		return
	}
	for {
		lc := getLfc(idx, inp)
		rc := getRtc(idx, inp)
		val := inp[idx]
		mp := minPoint(val, lc, rc)
		if isEqual(mp, val) {
			return
		} else if isEqual(mp, lc) {
			inp[idx], inp[(idx<<1)+1] = lc, val
			idx = (idx << 1) + 1
		} else {
			inp[idx], inp[(idx<<1)+2] = rc, val
			idx = (idx << 1) + 2
		}
	}
}

func minPoint(args ...[]int) []int {
	var mp []int
	for i := range args {
		if mp == nil || args[i] != nil && getED(args[i]) < getED(mp) {
			mp = args[i]
		}
	}
	return mp
}

func getED(inp []int) int {
	return inp[0]*inp[0] + inp[1]*inp[1]
}

func isEqual(inp1, inp2 []int) bool {
	return reflect.DeepEqual(inp1, inp2)
}

func createMinPointHeap(inp [][]int) {
	for i := (len(inp) - 2) >> 1; i >= 0; i-- {
		minPointHeapify(i, inp)
	}
}

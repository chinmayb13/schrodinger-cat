package heaps1

/*
Problem Description
Given a sorted array of integers A which contains 1 and some number of primes.
Then, for every p < q in the list, we consider the fraction p / q.

What is the B-th smallest fraction considered?

Return your answer as an array of integers, where answer[0] = p and answer[1] = q.

Problem Constraints
1 <= length(A) <= 2000
1 <= A[i] <= 30000
1 <= B <= length(A)*(length(A) - 1)/2

Input Format
The first argument of input contains the integer array, A.
The second argument of input contains an integer B.

Output Format
Return an array of two integers, where answer[0] = p and answer[1] = q.

Example Input
Input 1:
A = [1, 2, 3, 5]
B = 3

Input 2:
A = [1, 7]
B = 1

Example Output
Output 1:
[2, 5]

Output 2:
[1, 7]

Example Explanation
Explanation 1:
The fractions to be considered in sorted order are:
[1/5, 1/3, 2/5, 1/2, 3/5, 2/3]
The third fraction is 2/5.

Explanation 2:
The fractions to be considered in sorted order are:
[1/7]
The first fraction is 1/7.
*/
func GetKthSmallestFrac(A []int, B int) []int {
	var infos []*info
	for i := range A {
		for j := i + 1; j < len(A); j++ {
			infos = append(infos, getInfo(A[i], A[j]))
		}
	}
	//create MinHeap
	createMinInfoHeap(infos)
	for B > 1 {
		infos[0] = infos[len(infos)-1]
		infos = infos[:len(infos)-1]
		minInfoHeapify(0, infos)
		B--
	}
	//return the Bth smallest frac
	return []int{infos[0].num, infos[0].denom}
}

type info struct {
	num   int
	denom int
	frac  float64
}

func getInfoLC(idx int, lists []*info) *info {
	lc := (idx << 1) + 1
	if lc >= len(lists) {
		return nil
	}
	return lists[lc]
}

func getInfoRC(idx int, lists []*info) *info {
	rc := (idx << 1) + 2
	if rc >= len(lists) {
		return nil
	}
	return lists[rc]
}

func minInfoHeapify(idx int, lists []*info) {
	if len(lists) < 2 {
		return
	}

	for {
		lc := getInfoLC(idx, lists)
		rc := getInfoRC(idx, lists)
		cur := lists[idx]
		//if cur, lc and rc are nil, getMinInfo will return nil
		minInfo := getMinInfo(cur, lc, rc)
		if minInfo == cur {
			return
		} else if minInfo == lc {
			lists[idx], lists[(idx<<1)+1] = lc, cur
			idx = (idx << 1) + 1
		} else {
			lists[idx], lists[(idx<<1)+2] = rc, cur
			idx = (idx << 1) + 2
		}
	}
}

func getMinInfo(infos ...*info) *info {
	var minInfo *info
	for i := range infos {
		if minInfo == nil || (infos[i] != nil && (infos[i].frac < minInfo.frac)) {
			minInfo = infos[i]
		}
	}
	return minInfo
}

func createMinInfoHeap(lists []*info) {
	for i := (len(lists) - 2) >> 1; i >= 0; i-- {
		minInfoHeapify(i, lists)
	}
}

func getInfo(inp1, inp2 int) *info {
	return &info{
		num:   inp1,
		denom: inp2,
		frac:  float64(inp1) / float64(inp2),
	}
}

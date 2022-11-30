package subarray

//Two Pointer Approach:
func GetAlternateSubArrCentreIdxListAltII(A []int, B int) []int {
	if B == 0 {
		return getAllIndexes(len(A))
	}
	var indexArr []int
	subArrLength := 2*B + 1
	l, r := 0, 0
	for i := 1; i < len(A); i++ {
		//if alternating sequence, keep increasing right pointer
		if A[i] != A[i-1] {
			r++
			length := r - l + 1
			//if subarray is of needed length, store the centre and increment the left pointer
			if length == subArrLength {
				indexArr = append(indexArr, (l+r)/2)
				l++
			}
			//else start a fresh subarray, by resetting left and right ptr to current index
		} else {
			l, r = i, i
		}
	}
	return indexArr
}

func getAllIndexes(n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}
	return arr
}

func GetAlternateSubArrCentreIdxListAlt(A []int, B int) []int {
	var idxList []int
	subArraySize := 2*B + 1
	arrLength := len(A)
	prefixSum := make([]int, arrLength)
	for i := 0; i < arrLength; i++ {
		if i == 0 {
			prefixSum[i] = 1
			continue
		}
		//increase the prefixSum counter only when the current element is different from last element
		if A[i] != A[i-1] {
			prefixSum[i] = prefixSum[i-1] + 1
			//else reset the counter to 1
		} else {
			prefixSum[i] = 1
		}
	}

	for l, r := 0, subArraySize-1; r < arrLength; l, r = l+1, r+1 {
		//check if the sequence of 2xB+1 length is alternating or not
		if prefixSum[r] >= subArraySize {
			idxList = append(idxList, (l+r)/2)
		}
	}
	return idxList
}

package subarray

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

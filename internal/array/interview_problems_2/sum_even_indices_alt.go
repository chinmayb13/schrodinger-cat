package interviewproblems2

func GetSumEvenIndicesAlt(A []int, B [][]int) []int {
	var result []int
	queryLength := len(B)
	ps := buildEvenPrefixSum(A)
	for i := 0; i < queryLength; i++ {
		sum := 0
		startIdx := B[i][0]
		endIdx := B[i][1]
		if B[i][0] == 0 {
			sum += ps[endIdx]
		} else {
			sum += ps[endIdx] - ps[startIdx-1]
		}
		result = append(result, sum)
	}
	return result
}

func buildEvenPrefixSum(inp []int) []int {
	var prefixSum []int
	sum := 0
	for i := range inp {
		addendum := 0
		if i&1 == 0 {
			addendum = inp[i]
		}
		sum += addendum
		prefixSum = append(prefixSum, sum)
	}
	return prefixSum
}

package bureauid

/*
-2, 1, -3, 4, -1, 2, 1, -5, 4

max continuous subarray


*/

func GetMaxSumPossible(inp []int) int {
	if len(inp) == 0 {
		return 0
	}
	var maxSum int = -(1e9 + 7)
	sum := 0
	for i := range inp {
		if (sum + inp[i]) < 0 {
			sum = 0
			continue
		}
		sum = sum + inp[i]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum

}

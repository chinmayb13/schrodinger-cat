package backtracking1

func copyArr(inp []int) []int {
	var ans []int
	for i := range inp {
		ans = append(ans, inp[i])
	}
	return ans
}

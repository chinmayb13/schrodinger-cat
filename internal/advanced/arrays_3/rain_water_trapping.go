package arrays3

import "github.com/chinmayb13/schrodinger-cat/internal/helper"

func GetWaterCollected(A []int) int {
	totalWaterCollected := 0
	prefixMaxArr := buildPrefixMax(A)
	maxFromRight := 0
	for i := len(A) - 1; i >= 0; i-- {
		waterLevel := helper.Min(maxFromRight, prefixMaxArr[i])
		totalWaterCollected += helper.Max(waterLevel-A[i], 0)
		maxFromRight = helper.Max(maxFromRight, A[i])
	}
	return totalWaterCollected
}

func buildPrefixMax(inp []int) []int {
	prefixMaxArr := make([]int, len(inp))
	for i := 1; i < len(inp); i++ {
		prefixMaxArr[i] = helper.Max(inp[i-1], prefixMaxArr[i-1])
	}
	return prefixMaxArr
}

package primenumbers

import "math"

func countFactors(inputArr []int, inp int) int {
	ans := 1
	for inp > 1 {
		count := 0
		divisor := inputArr[inp]
		for (inp % divisor) == 0 {
			inp = inp / divisor
			count++
		}
		ans *= (count + 1)
	}
	return ans
}

func buildSieveArray(size int) []int {
	sieveArray := make([]int, size)
	for i := 0; i < size; i++ {
		sieveArray[i] = i
	}
	return sieveArray
}

func enrichSieveArray(sieveArray []int, inp int) {
	sqrt := int(math.Sqrt(float64(inp)))
	for i := 2; i <= sqrt; i++ {
		if sieveArray[i] == i {
			for j := i * i; j <= inp; j += i {
				if sieveArray[j] == j {
					sieveArray[j] = i
				}
			}
		}
	}
}

package interviewproblems2

func GetSpecialIndexCountAlt(A []int) int {
	idxCount := 0
	totalOddSum := 0
	totalEvenSum := 0
	leftEvenSum, leftOddSum := 0, 0
	for i := range A {
		if i&1 == 1 {
			totalOddSum += A[i]
		} else {
			totalEvenSum += A[i]
		}
	}

	for i := range A {
		if i&1 == 1 {
			rightEvenSum := totalOddSum - leftOddSum - A[i]
			rightOddSum := totalEvenSum - leftEvenSum

			evenSum := leftEvenSum + rightEvenSum
			oddSum := leftOddSum + rightOddSum
			if evenSum == oddSum {
				idxCount++
			}
			leftOddSum += A[i]
		} else {
			rightEvenSum := totalOddSum - leftOddSum
			rightOddSum := totalEvenSum - leftEvenSum - A[i]

			evenSum := leftEvenSum + rightEvenSum
			oddSum := leftOddSum + rightOddSum
			if evenSum == oddSum {
				idxCount++
			}
			leftEvenSum += A[i]
		}
	}
	return idxCount
}

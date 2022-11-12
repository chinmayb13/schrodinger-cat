package prefixsum

//approach: using two pointers starIdx and i, calculate min subarray of sum >=k
func MinSubArrayLen(target int, nums []int) int {
	startIdx := 0
	localSum := 0
	var minLength int = 1e5 + 1

	for i := range nums {
		localSum += nums[i]
		for localSum >= target && startIdx <= i {
			localLength := i - startIdx + 1
			if localLength < minLength {
				minLength = localLength
			}
			localSum = localSum - nums[startIdx]
			startIdx++
		}
	}

	if minLength == (1e5 + 1) {
		return 0
	}
	return minLength

}

func MinSubArrayLenAlt(target int, nums []int) int {
	var minLength int = 1e5 + 1

	for i := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target && (j-i+1) < minLength {
				minLength = j - i + 1
			}
		}
	}

	if minLength == (1e5 + 1) {
		return 0
	}
	return minLength

}

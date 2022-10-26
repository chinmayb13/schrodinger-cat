package prefixsum

func CheckSubarraySum(nums []int, k int) bool {
	var ps []int
	for i := range nums {
		if i == 0 {
			ps = append(ps, nums[i])
		} else {
			ps = append(ps, nums[i]+ps[i-1])
		}
	}
	rmnMap := map[int]int{ps[0] % k: 0}

	for i := 1; i < len(ps); i++ {

		remainder := ps[i] % k

		if remainder == 0 {
			return true
		}

		val, ok := rmnMap[remainder]

		if ok && i-val >= 2 {
			return true
		} else if !ok {
			rmnMap[ps[i]%k] = i
		}

	}
	return false
}

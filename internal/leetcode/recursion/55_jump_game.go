package recursion

func CanJump(nums []int) bool {
	reachableIdx := 0
	for i := range nums {
		if reachableIdx >= len(nums)-1 {
			return true
		} else if reachableIdx < i {
			return false
		}
		localReachableIdx := i + nums[i]
		if localReachableIdx > reachableIdx {
			reachableIdx = localReachableIdx
		}
	}
	return false
}

func isFeasible(idx int, nums []int) bool {
	if idx == len(nums)-1 {
		return true
	}
	if idx >= len(nums) {
		return false
	}

	for i := nums[idx]; i >= 1; i++ {
		check := isFeasible(idx+i, nums)
		if check {
			return true
		}
	}
	return false
}

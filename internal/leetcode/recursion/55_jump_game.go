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

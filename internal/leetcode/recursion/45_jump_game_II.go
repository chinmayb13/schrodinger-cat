package recursion

func Jump(nums []int) int {
	maxReachableIdx := 0
	currJump := 0
	jumps := 0
	for i := 0; i < len(nums)-1; i++ {
		localJump := i + nums[i]
		if maxReachableIdx < localJump {
			maxReachableIdx = localJump
		}
		if i == currJump {
			currJump = maxReachableIdx
			jumps++
		}
	}
	return jumps
}

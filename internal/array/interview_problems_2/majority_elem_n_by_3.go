package interviewproblems2

/*
Problem Description
You're given a read-only array of N integers. Find out if any integer occurs more than N/3 times in the array in linear time and constant additional space.
If so, return the integer. If not, return -1.

If there are multiple solutions, return any one.



Problem Constraints
1 <= N <= 7*105
1 <= A[i] <= 109


Input Format
The only argument is an integer array A.


Output Format
Return an integer.


Example Input
[1 2 3 1 1]


Example Output
1


Example Explanation
1 occurs 3 times which is more than 5/3 times.
*/
func GetThirdMajorityElement(nums []int) int {
	cme1 := 0
	cme2 := 0
	cme1Freq := 0
	cme2Freq := 0
	for i := range nums {
		switch {
		case cme1 == nums[i]:
			cme1Freq++
		case cme2 == nums[i]:
			cme2Freq++
		case cme1Freq == 0:
			cme1 = nums[i]
			cme1Freq = 1
		case cme2Freq == 0:
			cme2 = nums[i]
			cme2Freq = 1
		default:
			cme1Freq--
			cme2Freq--
		}
	}
	cme1Count := 0
	cme2Count := 0
	for i := range nums {
		if nums[i] == cme1 {
			cme1Count++
		}
		if nums[i] == cme2 {
			cme2Count++
		}
	}
	if cme1Count > len(nums)/3 {
		return cme1
	}
	if cme2Count > len(nums)/3 {
		return cme2
	}
	return -1
}

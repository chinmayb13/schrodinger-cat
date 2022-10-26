package bitmanipulations1

import (
	"math"
	"strconv"
)

//approach : Generate bitsequence from 0...0 to 2^n - 1 and against every bitmask value, add element to the subarray when bit value is set
func Subsets(nums []int) [][]int {
	var output [][]int
	arrLength := len(nums)
	shift := 1 << arrLength
	for i := 0; i < int(math.Pow(2, float64(arrLength))); i++ {
		b := strconv.FormatInt(int64(i)|int64(shift), 2)[1:]
		var subset []int
		for k := 0; k < len(b); k++ {
			if b[k] == '1' {
				subset = append(subset, nums[k])
			}
		}
		if subset == nil {
			subset = make([]int, 0)
		}
		output = append(output, subset)
	}
	return output
}

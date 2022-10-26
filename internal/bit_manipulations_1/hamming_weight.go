package bitmanipulations1

func HammingWeight(num uint32) int {
	countOne := 0
	var shift uint32 = 1
	for i := 0; i < 32; i++ {
		if shift&num > 0 {
			countOne++
		}
		shift <<= 1
	}
	return countOne

}

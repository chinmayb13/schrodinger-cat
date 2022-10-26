package carryforward

func GetSubSeqCountII(A string) int {
	//number of occurences of character A and X
	countA := 0
	countX := 0

	//Total subsequence count
	subSeqCount := 0
	for i := 0; i < len(A); i++ {
		if A[i] == 'A' {
			//increase character A count
			countA++

		} else if A[i] == 'X' {
			//increase character A count
			countX++
		} else if A[i] == 'G' {
			//the count would be as many AX combinations present when G was encountered + previous subsequence count
			subSeqCount += countA * countX
		}
	}
	return subSeqCount % (1e9 + 7)
}

package charstrings

/*
Problem Description
Given two binary strings, return their sum (also a binary string).

Example:
a = "100"
b = "11"
Return a + b = "111".
*/
func AddBinary(A string, B string) string {
	carry := 0
	var addition []byte
	idx_A := len(A) - 1
	idx_B := len(B) - 1
	for idx_A >= 0 || idx_B >= 0 || carry > 0 {
		var bit_A, bit_B int
		if idx_A >= 0 {
			bit_A = int(A[idx_A] - '0')
		}
		if idx_B >= 0 {
			bit_B = int(B[idx_B] - '0')
		}
		sum := bit_A + bit_B + carry
		addition = append(addition, byte(sum%2)+'0')
		carry = sum / 2
		idx_A--
		idx_B--
	}

	return reverse(addition)
}

func reverse(runeArr []byte) string {
	i := 0
	j := len(runeArr) - 1
	for i < j {
		runeArr[i], runeArr[j] = runeArr[j], runeArr[i]
		i++
		j--
	}
	return string(runeArr)
}

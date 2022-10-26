package bitmanipulations1

import (
	"fmt"
	"strings"
)

func AddBinary(a string, b string) string {
	var sb strings.Builder
	if len(a) < len(b) {
		return AddBinary(b, a)
	}
	j := len(b) - 1
	carry := 0
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '1' {
			carry++
		}
		if j >= 0 && b[j] == '1' {
			carry++
		}

		sb.WriteString(fmt.Sprint(carry % 2))
		carry /= 2
		j--
	}
	if carry == 1 {
		sb.WriteString("1")
	}
	return reverseString(sb.String())
}

func reverseString(s string) string {
	var rev strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		rev.WriteByte(s[i])
	}
	return rev.String()
}

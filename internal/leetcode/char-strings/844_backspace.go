package charstrings

import "strings"

func BackspaceCompare(s string, t string) bool {
	return cleanArray(s) == cleanArray(t)
}

func cleanArray(s string) string {
	var sb strings.Builder
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			count++
			continue
		} else if count > 0 {
			count--
			continue
		} else {
			sb.WriteByte(s[i])
		}
	}
	return sb.String()
}

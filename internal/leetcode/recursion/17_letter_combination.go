package recursion

func LetterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	return getCombination("", 0, digits)
}

func getCombination(input string, idx int, digits string) []string {
	if idx == len(digits) {
		return []string{input}
	}
	var combination []string
	for ci := range numPad[digits[idx]] {
		localInput := input + string(numPad[digits[idx]][ci])
		subCombination := getCombination(localInput, idx+1, digits)
		if subCombination != nil {
			combination = append(combination, subCombination...)
		}
	}
	return combination
}

var numPad = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

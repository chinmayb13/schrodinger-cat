package gcd

func getGCD(inp1, inp2 int) int {
	if inp2 == 0 {
		return inp1
	}
	return getGCD(inp2, inp1%inp2)
}

package helper

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Abs(inp int) int {
	if inp < 0 {
		return inp * -1
	}
	return inp
}

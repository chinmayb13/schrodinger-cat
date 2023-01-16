package binarysearch3

/*
Problem Description
You are given three positive integers, A, B, and C.
Any positive integer is magical if divisible by either B or C.
Return the Ath smallest magical number. Since the answer may be very large, return modulo 109 + 7.

Problem Constraints
1 <= A <= 109
2 <= B, C <= 40000

Input Format
The first argument given is an integer A.
The second argument given is an integer B.
The third argument given is an integer C.

Output Format
Return the Ath smallest magical number. Since the answer may be very large, return modulo 109 + 7.

Example Input
Input 1:
A = 1
B = 2
C = 3

Input 2:
A = 4
B = 2
C = 3

Example Output
Output 1:
2

Output 2:
6

Example Explanation
Explanation 1:
1st magical number is 2.

Explanation 2:
First four magical numbers are 2, 3, 4, 6 so the 4th magical number is 6.
*/
func GetAthMagicalNum(A int, B int, C int) int {
	var mod int = 1e9 + 7
	var ans int
	lcm := (B * C) / getGCD(B, C)
	minimum := min(B, C)
	low, high := minimum, minimum*A
	for low <= high {
		mid := low + (high-low)>>1
		if ((mid / B) + (mid / C) - (mid / lcm)) >= A {
			ans = mid % mod
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}

func getGCD(inp1, inp2 int) int {
	if inp2 == 0 {
		return inp1
	}
	return getGCD(inp2, inp1%inp2)
}

func min(inp1, inp2 int) int {
	if inp1 < inp2 {
		return inp1
	}
	return inp2
}

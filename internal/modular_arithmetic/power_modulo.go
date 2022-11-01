package modulararithmetic

/*
Problem Description
You are given A, B and C .
Calculate the value of (A ^ B) % C


Problem Constraints
1 <= A <= 109
0 <= B <= 105
1 <= C <= 109


Input Format
Given three integers A, B and C.


Output Format
Return an integer.


Example Input
Input 1:
A = 2
B = 3
C = 3

Input 2:
A = 5
B = 2
C = 4


Example Output

Output 1:
2

Output 2:
1

Example Explanation

For Input 1:
(2 ^ 3) % 3 = 8 % 3 = 2

For Input 2:
(5 ^ 2) % 4 = 25 % 4 = 1
*/

func GetPowerMod(A int, B int, C int) int {
	num := 1
	A = A % C
	for B > 0 {
		if B&1 == 1 {
			num = (num * A) % C
			B = B - 1
		} else {
			A = (A * A) % C
			B = B / 2
		}
	}
	return num
}

func GetPowerModAlt(A int, B int, C int) int {
	if B == 0 {
		return 1 % C
	}
	if B&1 == 1 {
		return ((A % C) * GetPowerModAlt(((A%C)*(A%C))%C, (B-1)/2, C)) % C
	}
	return (GetPowerModAlt(((A%C)*(A%C))%C, B/2, C)) % C
}

// func GetPowerModAlternate(A int, B int, C int) int {
// 	var num int64 = 1
// 	for i := 0; i < B; i++ {
// 		num = (num * int64(A%C)) % int64(C)
// 	}
// 	return int(num)
// }

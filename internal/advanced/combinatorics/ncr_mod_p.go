package combinatorics

/*
Problem Description
Given three integers A, B, and C, where A represents n, B represents r, and C represents p and p is a prime number greater than equal to n, find and return the value of nCr % p where nCr % p = (n! / ((n-r)! * r!)) % p.
x! means factorial of x i.e. x! = 1 * 2 * 3... * x.
NOTE: For this problem, we are considering 1 as a prime.

Problem Constraints
1 <= A <= 106
1 <= B <= A
A <= C <= 109+7

Input Format
The first argument given is the integer A ( = n).
The second argument given is the integer B ( = r).
The third argument given is the integer C ( = p).

Output Format
Return the value of nCr % p.

Example Input
Input 1:
A = 5
B = 2
C = 13

Input 2:
A = 6
B = 2
C = 13

Example Output
Output 1:
10

Output 2:
2

Example Explanation
Explanation 1:
nCr( n=5 and r=2) = 10.
p=13. Therefore, nCr%p = 10.
*/
func GetnCrModP(A int, B int, C int) int {
	return (calcFact(A, C) * calcPower(int64(calcFact(A-B, C)*calcFact(B, C))%int64(C), C-2, C)) % C
}

func calcFact(inp, mod int) int {
	var fact int64 = 1
	for i := 1; i <= inp; i++ {
		fact = (fact * int64(i)) % int64(mod)
	}
	return int(fact)
}

func calcPower(inp int64, pow, mod int) int {
	var ans int64 = 1
	for pow > 0 {
		if (pow & 1) > 0 {
			ans = (ans * inp) % int64(mod)
			pow--
		} else {
			inp = (inp * inp) % int64(mod)
			pow = pow >> 1
		}
	}
	return int(ans)
}

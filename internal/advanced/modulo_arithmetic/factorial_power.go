package moduloarithmetic

/*
Problem Description
Given two Integers A, B. You have to calculate (A ^ (B!)) % (1e9 + 7).
"^" means power,
"%" means "mod", and
"!" means factorial.

Problem Constraints
1 <= A, B <= 5e5

Input Format
First argument is the integer A
Second argument is the integer B

Output Format
Return one integer, the answer to the problem

Example Input
Input 1:
A = 1
B = 1

Input 2:
A = 2
B = 2

Example Output
Output 1:
1

Output 2:
4

Example Explanation
Explanation 1:
1! = 1. Hence 1^1 = 1.

Explanation 2:
2! = 2. Hence 2^2 = 4.
*/

/*
   Solution Approach:

   We have to find A ^ B! now as we can notice for larger values
   of B! the result num gets too large fairly quickly and could also cause overflow
   so to optimimize this we need to find a optimal way to calculate B!

   As we can notice from the problem statement we need to return A ^ B! % 1e9 + 7
   Here 1e9 + 7 is a prime number. This looks somewhat similar to Fermat’s Little Theorem, i.e,

   → A ^ (P-1) % P = 1 (1)
   we can write A^ B! as;

   A ^ B! = A ^ (p-1) * A ^ (p-1) * A ^ (p-1) … A ^ X  (2)

   But since we need to find A^B! % P; we can write eqn 2 as
   (A ^ B!) % P = (A ^ (p-1) * A ^ (p-1) * A ^ (p-1) … A ^ X) % P  ___(3)

   Using assiociative property of modulo multiplication;
   → (ab)%p = ((a % P)(b % P)) % P

   So every term with A ^ (p-1) in eqn 3 will be == 1 by applying FLT [eqn(1)]
   and we will be only left with A^X % P on RHS in eqn 3
   → A^B! % P = A^X % P

   Now X is nothing but B! % (P - 1) since we continuously substracted P-1 from B!
   And as we know repetative substraction is division.

   So our problem has now been reduced down to finding
   1. value of; X = B! % (P - 1)
   2. value of; A ^ X

   finally we can return ans % P
*/
func GetLargePower(A int, B int) int {
	var mod int = 1e9 + 7
	return power(int64(A), factorial(B, mod), mod)

}

func factorial(inp int, mod int) int {
	var fact int64 = 1
	for i := 1; i <= inp; i++ {
		fact = (fact * int64(i)) % int64(mod-1)
	}
	return int(fact)
}

func power(inp int64, pow, mod int) int {
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

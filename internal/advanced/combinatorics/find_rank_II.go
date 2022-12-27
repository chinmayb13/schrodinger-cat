package combinatorics

/*
Problem Description
Given a string A, find the rank of the string amongst its permutations sorted lexicographically. Note that the characters might be repeated. If the characters are repeated, we need to look at the rank in unique permutations. Look at the example for more details.

NOTE:
The answer might not fit in an integer, so return your answer % 1000003 where 1000003 is a prime number.
String A can consist of both lowercase and uppercase letters. Characters with lesser ASCII values are considered smaller, i.e., 'a' > 'Z'.

Problem Constraints
1 <= len(A) <= 1000000

Input Format
First argument is a string A.

Output Format
Return an integer denoting the rank.

Example Input
Input 1:
A = "aba"

Input 2:
A = "bca"

Example Output
Output 1:
2

Output 2:
4

Example Explanation
Explanation 1:

The order permutations with letters 'a', 'a', and 'b' :
   aab
   aba
   baa
So, the rank is 2.

Explanation 2:

The order permutations with letters 'a', 'b', and 'c' :
   abc
   acb
   bac
   bca
   cab
   cba
So, the rank is 4.
*/
func FindRankII(A string) int {
	mod := 1000003
	ans := 0
	for i := 0; i < len(A)-1; i++ {
		//map for frequency
		hashMap := make(map[byte]int)
		//map containing element smaller than current element
		smallerSet := make(map[byte]struct{})
		divisor := 1
		hashMap[A[i]]++
		for j := i + 1; j < len(A); j++ {
			if A[j] < A[i] {
				smallerSet[A[j]] = struct{}{}
			}
			hashMap[A[j]]++
			//divisor calculating factorial of repeating elements
			divisor = (divisor * hashMap[A[j]]) % mod
		}
		numerator := fact(len(A)-i-1, mod)
		for k := range smallerSet {
			val := hashMap[k] - 1
			localDivisor := ((divisor * calcPower(int64(fact(hashMap[k], mod)), mod-2, mod)) % mod * fact(val, mod)) % mod
			ans = (ans + (numerator*calcPower(int64(localDivisor), mod-2, mod))%mod) % mod

		}

	}
	return (ans + 1) % mod
}

func fact(inp, mod int) int {
	var fact int64 = 1
	for i := 2; i <= inp; i++ {
		fact = (fact * int64(i)) % int64(mod)
	}
	return int(fact)
}

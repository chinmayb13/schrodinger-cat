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
	ans := 0
	mod := 1000003
	/*get frequency array of size 52 (a-z A-Z)

	  divisor contains the product of factorial of the frequency of all elements
	  for a string: bbbaaacccc, divisor will contain 3!*3!*4!*/
	freqArray, divisor := getFreqArray(A)
	//get factorial array
	factArray := getFactorialArray(len(A) + 1)
	for i := range A {
		idx := getIdx(A[i])
		numerator := len(A) - 1 - i
		for j := 0; j < idx; j++ {
			/*if there exists a lexicographical smaller element than the current ith element
			  place that smaller element in the current position and count the permutation
			  permutatation = factorial of length of remaining string after putting jth element in current postion /
			                  (divisor/factorial of frequency of jth element)*(factorial of one less than the frequency of jth element)

			*/
			if freqArray[j] > 0 {
				localDivisor := (((divisor * inverse(factArray[freqArray[j]])) % mod) * factArray[freqArray[j]-1]) % mod
				ans = (ans + (factArray[numerator]*inverse(localDivisor))%mod) % mod
			}
		}
		/*
			since we have covered the current element's position, we will decrease the frequency of current element by 1
			and update the divisor as
			(divisor/factorial of frequency of ith element)*(factorial of one less than the frequency of ith element)
		*/
		divisor = (((divisor * inverse(factArray[freqArray[idx]])) % mod) * factArray[freqArray[idx]-1]) % mod
		freqArray[idx]--
	}
	return ans + 1
}

func getFreqArray(inp string) ([]int, int) {
	freqArray := make([]int, 52)
	divisor := 1
	mod := 1000003

	for i := range inp {
		idx := getIdx(inp[i])
		freqArray[idx]++
		divisor = (divisor * freqArray[idx]) % mod
	}
	return freqArray, divisor
}

func getFactorialArray(len int) []int {
	mod := 1000003
	factArray := make([]int, len)
	factArray[0] = 1
	for i := 1; i < len; i++ {
		factArray[i] = (i * factArray[i-1]) % mod
	}
	return factArray
}

func inverse(inp int) int {
	mod := 1000003
	return calcPower(int64(inp), mod-2, mod)
}

func getIdx(inp byte) int {
	if inp >= 'a' {
		return int(inp) - int('a') + 26
	}
	return int(inp) - int('A')

}

func FindRankIIAlt(A string) int {
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

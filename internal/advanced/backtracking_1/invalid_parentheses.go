package backtracking1

import "sort"

/*
Problem Description
Given a string A consisting of lowercase English alphabets and parentheses '(' and ')'. Remove the minimum number of invalid parentheses in order to make the input string valid.

Return all possible results in sorted order.


Problem Constraints
1 <= length of the string <= 20

Input Format
The only argument given is string A.

Output Format
Return all possible strings after removing the minimum number of invalid parentheses.

Example Input
Input 1:
A = "()())()"

Input 2:
A = "(a)())()"

Example Output
Output 1:
["()()()", "(())()"]

Output 2:
["(a)()()", "(a())()"]

Example Explanation
Explanation 1:
By removing 1 parentheses we can make the string valid.
       1. Remove the parentheses at index 4 then string becomes : "()()()"
       2. Remove the parentheses at index 2 then string becomes : "(())()"

Explanation 2:
By removing 1 parentheses we can make the string valid.
       1. Remove the parentheses at index 5 then string becomes : "(a)()()"
       2. Remove the parentheses at index 2 then string becomes : "(a())()"
*/

func GetValidExpr(A string) []string {
	inp := A
	var invalidL, invalidR int
	//count number of invalid left and right parentheses
	for i := range inp {
		if inp[i] == '(' {
			invalidL++
		} else if inp[i] == ')' {
			if invalidL == 0 {
				invalidR++
			} else {
				invalidL--
			}
		}
	}

	//hashSet to avoid duplicate expressions
	hashSet := make(map[string]bool)
	var ans []string
	generateValidExpr(inp, 0, "", &ans, hashSet, invalidL, invalidR, 0)
	sort.Strings(ans)
	return ans
}

func generateValidExpr(inp string, idx int, cur string, ans *[]string, hashSet map[string]bool, invalidL, invalidR, balance int) {

	//balance: difference between count of left and right parentheses
	//if balance or invalid right or left parentheses are negative, return
	if balance < 0 || invalidL < 0 || invalidR < 0 {
		return
	}

	//if end of string, and everything is balanced, save the ans
	if idx == len(inp) {
		if balance+invalidL+invalidR == 0 && !hashSet[cur] {
			*ans = append(*ans, cur)
			hashSet[cur] = true
		}
		return
	}
	if inp[idx] == '(' {
		//don't include the parentheses, decrease the invalid left parentheses count
		generateValidExpr(inp, idx+1, cur, ans, hashSet, invalidL-1, invalidR, balance)
		//include the parentheses, increase the balance count
		generateValidExpr(inp, idx+1, cur+"(", ans, hashSet, invalidL, invalidR, balance+1)
	} else if inp[idx] == ')' {
		//don't include the parentheses, decrease the invalid left parentheses count
		generateValidExpr(inp, idx+1, cur, ans, hashSet, invalidL, invalidR-1, balance)
		//include the parentheses, decrease the balance count
		generateValidExpr(inp, idx+1, cur+")", ans, hashSet, invalidL, invalidR, balance-1)
	} else {
		//move on
		generateValidExpr(inp, idx+1, cur+string(inp[idx]), ans, hashSet, invalidL, invalidR, balance)
	}
}

package dp6

/*
Problem Description
Given a string A made up of multiple brackets of type "[]" , "()" and "{}" find the length of the longest substring which forms a balanced string .

Conditions for a string to be balanced :

Blank string is balanced ( However blank string will not be provided as a test case ).
If B is balanced : (B) , [B] and {B} are also balanced.
If B1 and B2 are balanced then B1B2 i.e the string formed by concatenating B1 and B2 is also balanced.

Problem Constraints
0 <= |A| <= 106

Input Format
First and only argument is an string A.

Output Format
Return an single integer denoting the lenght of the longest balanced substring.

Example Input
Input 1:
A = "[()]"

Input 2:
A = "[(])"

Example Output
Output 1:
4

Output 2:
0

Example Explanation
Explanation 1:
Substring [1:4] i.e "[()]" is the longest balanced substring of length 4.

Explanation 2:
None of the substring is balanced so answer is 0.
*/
func LBSlength(A string) int {
	inp := A
	N := len(inp)
	memo := make([]int, N+1)

	pMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	for i := 1; i <= N; i++ {
		//check for closing braces
		if i > 1 && (inp[i-1] == ')' || inp[i-1] == ']' || inp[i-1] == '}') {
			match := pMap[inp[i-1]]
			//if the last character was open brace
			//add 2 to the answer at index before open brace
			if inp[i-2] == match {
				memo[i] = memo[i-2] + 2
			}

			//if there was valid parentheses ending at last character
			//check if there is an open brace before the valid parentheses length
			//if it is, save ans before the open brace + valid parentheses length + 2

			if memo[i-1] > 0 && i-memo[i-1]-2 >= 0 && inp[i-memo[i-1]-2] == match {
				memo[i] = memo[i-memo[i-1]-2] + memo[i-1] + 2
			}
		}
	}

	ans := 0
	for i := 1; i <= N; i++ {
		ans = max(ans, memo[i])
	}
	return ans
}

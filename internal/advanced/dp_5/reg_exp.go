package dp5

/*
Problem Description
Implement wildcard pattern matching with support for ' ? ' and ' * ' for strings A and B.

' ? ' : Matches any single character.
' * ' : Matches any sequence of characters (including the empty sequence).
The matching should cover the entire input string (not partial).

Problem Constraints
1 <= length(A), length(B) <= 104

Input Format
The first argument of input contains a string A.
The second argument of input contains a string B.

Output Format
Return 1 if the patterns match else return 0.

Example Input
Input 1:
A = "aaa"
B = "a*"

Input 2:
A = "acz"
B = "a?a"

Example Output
Output 1:
1

Output 2:
0

Example Explanation
Explanation 1:
Since '*' matches any sequence of characters. Last two 'a' in string A will be match by '*'.
So, the pattern matches we return 1.

Explanation 2:
'?' matches any single character. First two character in string A will be match.
But the last character i.e 'z' != 'a'. Return 0.
*/
func IsMatch(A string, B string) int {
	source := A
	target := B
	N := len(source)
	M := len(target)
	memo := make([][]bool, N+1)
	for i := range memo {
		memo[i] = make([]bool, M+1)
	}

	for i := range target {
		if B[i] == '*' {
			//if * is encountered at ith index of the target
			//with index i+1 to N in the source, and index i+1 in the target
			//set value to true
			for j := i + 1; j <= N; j++ {
				memo[j][i+1] = true
			}
			//if we have a series of * from 0th to ith index
			//with index 0 in the source, and 1 to i+1 in the target
			//set value to true
			if i == 0 {
				for k := i; k < len(target); k++ {
					if target[k] != '*' {
						break
					}
					memo[0][k+1] = true
				}
			}
		}
	}

	memo[0][0] = true

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			//take OR of keeping input character and excluding *
			//and excluding the input character and including *
			if target[j-1] == '*' {
				memo[i][j] = memo[i-1][j] || memo[i][j-1]
				//else check the previous indexes of source and input
			} else if target[j-1] == '?' || target[j-1] == source[i-1] {
				memo[i][j] = memo[i-1][j-1]
			}
		}
	}
	if memo[N][M] {
		return 1
	}
	return 0
}

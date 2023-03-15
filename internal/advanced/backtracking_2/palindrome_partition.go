package backtracking2

/*
Problem Description
Given a string A, partition A such that every string of the partition is a palindrome.

Return all possible palindrome partitioning of A.

Ordering the results in the answer : Entry i will come before Entry j if :
len(Entryi[0]) < len(Entryj[0]) OR
(len(Entryi[0]) == len(Entryj[0]) AND len(Entryi[1]) < len(Entryj[1])) OR * * *
(len(Entryi[0]) == len(Entryj[0]) AND ... len(Entryi[k] < len(Entryj[k]))

Problem Constraints
1 <= len(A) <= 15

Input Format
First argument is a string A of lowercase characters.

Output Format
Return a list of all possible palindrome partitioning of s.

Example Input
Input 1:
A = "aab"

Input 2:
A = "a"

Example Output
Output 1:
[
   ["a","a","b"]
   ["aa","b"],
 ]

Output 2:
[
   ["a"]
 ]

Example Explanation
Explanation 1:
In the given example, ["a", "a", "b"] comes before ["aa", "b"] because len("a") < len("aa").

Explanation 2:
In the given example, only partition possible is "a" .
*/
func Partition(s string) [][]string {
	var ans [][]string
	var curArr []string
	getPartitions(0, s, curArr, &ans)
	return ans
}

func getPartitions(idx int, inp string, curArr []string, ans *[][]string) {
	//if end of inp has reached, save the array to the ans
	if idx == len(inp) {
		*ans = append(*ans, copyStrArr(curArr))
		return
	}

	//take empty string
	var cur string

	//iterate from idx to end
	for i := idx; i < len(inp); i++ {
		//add current byte to the string
		cur = cur + string(inp[i])
		//if the string is palindrome, add the string to the array, check for the next idx
		if isPalindrom(cur) {
			curArr = append(curArr, cur)
			getPartitions(i+1, inp, curArr, ans)
			//remove the last string
			curArr = curArr[:len(curArr)-1]
		}
	}
}

func isPalindrom(inp string) bool {
	for i, j := 0, len(inp)-1; i < j; i, j = i+1, j-1 {
		if inp[i] != inp[j] {
			return false
		}
	}
	return true
}

func copyStrArr(inp []string) []string {
	var ans []string
	for i := range inp {
		ans = append(ans, inp[i])
	}
	return ans
}

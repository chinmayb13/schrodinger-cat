package recursion2

/*
Problem Description
Given the total number of person A and a number B which indicates that B-1 persons are skipped and the Bth person is killed in a circle. Find the last person standing in the circle.

Problem Constraints
1 <= A <= 105
2 <= B <= A

Input Format
First argument A is an integer.
Second argument B is an integer.

Output Format
Return an integer.

Example Input
Input 1:
A = 4
B = 2

Input 2:
A = 5
B = 3

Example Output
Output 1:
1

Output 2:
4

Example Explanation
For Input 1:
Firstly, the person at position 2 is killed, then the person at position 4 is killed,
then the person at position 3 is killed. So the person at position 1 survives.

For Input 2:
Firstly, the person at position 3, then the person at position 1 is killed,
then the person at position 5 is killed. Finally, the person at position 2 is killed.
So the person at position 4 survives.
*/
func GetLastPerson(A int, B int) int {
	arr := make([]int, A)
	for i := 0; i < A; i++ {
		arr[i] = i + 1
	}
	val := getLastPersonStanding(arr, B-1)
	return val[0]
}

func getLastPersonStanding(arr []int, B int) []int {
	if len(arr) == 1 {
		return arr
	}

	idx := B % len(arr)
	return getLastPersonStanding(append(arr[idx+1:], arr[:idx]...), B)

}

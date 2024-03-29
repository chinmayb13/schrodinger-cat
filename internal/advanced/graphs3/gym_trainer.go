package graphs3

/*
Problem Description

You are the trainer of a gym and there are A people who come to your gym.
Some of them are friends because they walk together, and some of them are friends because they talk together.
But people become possessive about each other, so a person cannot walk with one friend and talk with another. Although he can walk with two or more people or talk with two or more people.
You being the trainer, decided to suggest each one of the 2 possible diets. But friends, being friends will always have the same diet as all the other friends are having.
Find and return the number of ways you can suggest each of them a diet.
As the number of ways can be huge, return the answer modulo 109+7.

NOTE: If there is any person who walks with one person and talks with the another person, then you may return 0.

Problem Constraints
1 <= A, N, M <= 105
1 <= B[i][0], B[i][1], C[i][0], C[i][1] <= A

Input Format
The first argument contains an integer A, representing the number of people.
The second argument contains a 2-D integer array B of size N x 2, where B[i][0] and B[i][1] are friends because they walk together.
The third argument contains a 2-D integer array C of size M x 2, where C[i][0] and C[i][1] are friends because they talk together.

Output Format
Return an integer representing the number of ways to suggest one of the two diets to the people.

Example Input
Input 1:
A = 4
B = [

	  [1, 2]
	]

C = [

	  [3, 4]
	]

Input 2:
A = 3
B = [

	  [1, 2]
	]

C = [

	  [1, 3]
	]

Example Output
Output 1:
4

Output 2:
0

# Example Explanation

Explanation 1:
There are four ways to suggest the diet:
Diet-1 to (1, 2) and Diet-2 to (3, 4).
Diet-1 to (1, 2) and Diet-1 to (3, 4).
Diet-2 to (1, 2) and Diet-1 to (3, 4).
Diet-2 to (1, 2) and Diet-2 to (3, 4).

Explanation 2:
Person 1 walks with person 2 and talks with person 3. So, we will return 0.
*/
func GetDietWays(A int, B [][]int, C [][]int) int {
	var mod int = 1e9 + 7
	adjListWalk := make([][]int, A+1)
	//create adjacency list for walking friends
	for i := range B {
		adjListWalk[B[i][0]] = append(adjListWalk[B[i][0]], B[i][1])
		adjListWalk[B[i][1]] = append(adjListWalk[B[i][1]], B[i][0])
	}

	adjListTalk := make([][]int, A+1)
	//create adjacency list for talking friends
	for i := range C {
		if len(adjListWalk[C[i][0]]) > 0 || len(adjListWalk[C[i][1]]) > 0 {
			return 0
		}
		adjListTalk[C[i][0]] = append(adjListTalk[C[i][0]], C[i][1])
		adjListTalk[C[i][1]] = append(adjListTalk[C[i][1]], C[i][0])
	}

	//visited array
	visited := make([]bool, A+1)
	ans := 1
	for i := 1; i <= A; i++ {
		//if not visited, do DFS
		// since diet has 2 options, each connected component can have 2 values
		//hence ans = (num of CC)^2
		if !visited[i] {
			ans = (ans << 1) % mod
			dfs(i, adjListWalk, visited)
			dfs(i, adjListTalk, visited)
		}
	}
	return ans
}

func dfs(vertice int, adjList [][]int, visited []bool) {
	visited[vertice] = true
	for i := range adjList[vertice] {
		v := adjList[vertice][i]
		if !visited[v] {
			dfs(v, adjList, visited)
		}
	}
}

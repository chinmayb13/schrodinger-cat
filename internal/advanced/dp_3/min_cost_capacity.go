package dp3

import "math"

/*
Problem Description
As it is Tushar's Birthday on March 1st, he decided to throw a party to all his friends at TGI Fridays in Pune. Given are the eating capacity of each friend, filling capacity of each dish and cost of each dish. A friend is satisfied if the sum of the filling capacity of dishes he ate is equal to his capacity. Find the minimum cost such that all of Tushar's friends are satisfied (reached their eating capacity).

NOTE:

Each dish is supposed to be eaten by only one person. Sharing is not allowed.
Each friend can take any dish unlimited number of times.
There always exists a dish with filling capacity 1 so that a solution always exists.

Problem Constraints
|A| <= 1000
|B| <= 1000
|C| <= 1000

Input Format
First Argument is vector A, denoting eating capacities
Second Argument is vector B, denoting filling capacities
Third Argument is vector C, denoting cost

Output Format
Return a single integer, the answer to the problem

Example Input
Input 1:
A = [2, 4, 6]
B = [2, 1, 3]
C = [2, 5, 3]

Input 2:
A = [2]
B = [1]
C = [2]

Example Output
Output 1:
12

Output 2:
4

Example Explanation
Explanation 1:
First friend takes dish 1, Second friend takes dish 1 twice and third friend takes dish 3 twice.
So 2 + 2*2 + 3*2 = 12.

Explanation 2:
Only way is to take 2 dishes of cost 2, hence 4.
*/
func GetMinCostCapacity(A []int, B []int, C []int) int {
	maxCap := 0
	//find max capacity
	for i := range A {
		maxCap = max(maxCap, A[i])
	}

	//create array till max capacity
	memo := make([]int, maxCap+1)

	for i := 1; i <= maxCap; i++ {
		minCost := math.MaxInt
		for j := 0; j < len(B); j++ {
			//if current capacity is greater than dish capacity
			if i >= B[j] {
				//take min of including the dish vs excluding the dish
				minCost = min(minCost, memo[i-B[j]]+C[j])
			}
		}
		//add minCost for given capacity
		memo[i] = minCost
	}
	ans := 0

	//iterate over friends' capacity and add the cost
	for i := range A {
		ans += memo[A[i]]
	}
	return ans

}

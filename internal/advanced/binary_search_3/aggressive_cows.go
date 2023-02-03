package binarysearch3

import "sort"

/*
Problem Description
Farmer John has built a new long barn with N stalls. Given an array of integers A of size N where each element of the array represents the location of the stall and an integer B which represents the number of cows.
His cows don't like this barn layout and become aggressive towards each other once put into a stall. To prevent the cows from hurting each other, John wants to assign the cows to the stalls, such that the minimum distance between any two of them is as large as possible. What is the largest minimum distance?

Problem Constraints
2 <= N <= 100000
0 <= A[i] <= 109
2 <= B <= N

Input Format
The first argument given is the integer array A.
The second argument given is the integer B.

Output Format
Return the largest minimum distance possible among the cows.

Example Input
Input 1:
A = [1, 2, 3, 4, 5]
B = 3

Input 2:
A = [1, 2]
B = 2

Example Output
Output 1:
2

Output 2:
1

Example Explanation
Explanation 1:
John can assign the stalls at location 1, 3 and 5 to the 3 cows respectively. So the minimum distance will be 2.

Explanation 2:
The minimum distance will be 1.
*/
func GetMaxMinStallDist(A []int, B int) int {
	sort.Ints(A)
	ans := 0
	low := 1
	high := A[len(A)-1]
	for low <= high {
		mid := low + (high-low)>>1
		dist := getMinDistWithCowsTied(A, B-1, mid)
		if dist >= 0 {
			ans = dist
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans

}

func getMinDistWithCowsTied(stalls []int, cowsLeft int, minDistancePerCow int) int {
	source := 0
	var leastDistance int = 1e9 + 1
	for cowsLeft > 0 {
		lastSource := source
		searchItem := stalls[source] + minDistancePerCow
		low := source + 1
		high := len(stalls) - 1
		for low <= high {
			mid := low + (high-low)>>1
			if stalls[mid] >= searchItem {
				source = mid
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		if source == lastSource {
			return -1
		}
		cowsLeft--
		dist := stalls[source] - stalls[lastSource]
		if dist < leastDistance {
			leastDistance = dist
		}
	}
	return leastDistance
}
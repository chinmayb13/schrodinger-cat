package greedyalgo

import (
	"container/heap"
	"sort"
)

/*
Problem Description
Given two arrays, A and B of size N. A[i] represents the time by which you can buy the ith car without paying any money.
B[i] represents the profit you can earn by buying the ith car. It takes 1 minute to buy a car, so you can only buy the ith car when the current time <= A[i] - 1.

Your task is to find the maximum profit one can earn by buying cars considering that you can only buy one car at a time.
NOTE:
You can start buying from time = 0.
Return your answer modulo 109 + 7.

Problem Constraints
1 <= N <= 105
1 <= A[i] <= 109
0 <= B[i] <= 109

Input Format
The first argument is an integer array A represents the deadline for buying the cars.
The second argument is an integer array B represents the profit obtained after buying the cars.

Output Format
Return an integer denoting the maximum profit you can earn.

Example Input
Input 1:
A = [1, 3, 2, 3, 3]
B = [5, 6, 1, 3, 9]

Input 2:
A = [3, 8, 7, 5]
B = [3, 1, 7, 19]

Example Output
Output 1:
20

Output 2:
30

Example Explanation
Explanation 1:
At time 0, buy car with profit 5.
At time 1, buy car with profit 6.
At time 2, buy car with profit 9.
At time = 3 or after , you can't buy any car, as there is no car with deadline >= 4.
So, total profit that one can earn is 20.

Explanation 2:
At time 0, buy car with profit 3.
At time 1, buy car with profit 1.
At time 2, buy car with profit 7.
At time 3, buy car with profit 19.
We are able to buy all cars within their deadline. So, total profit that one can earn is 30.
*/
func GetMaxProfit(A []int, B []int) int {
	var mod int = 1e9 + 7
	profit := 0
	N := len(A)
	carInfos := make([]carInfo, N)
	for i := 0; i < N; i++ {
		carInfos[i].deadline = A[i]
		carInfos[i].profit = B[i]
	}

	//sort by the deadline to get the chronological order
	sort.Slice(carInfos, func(i, j int) bool {
		return carInfos[i].deadline < carInfos[j].deadline
	})

	mh := make(minIntHeap, 0)
	heap.Init(&mh)

	time := 0
	for i := 0; i < N; i++ {
		//if deadline is not missed
		if carInfos[i].deadline > time {
			heap.Push(&mh, carInfos[i].profit)
			time++
			//if deadline is missed, but the profit is promising
		} else if mh[0] < carInfos[i].profit {
			heap.Pop(&mh)
			heap.Push(&mh, carInfos[i].profit)
		}
	}

	for len(mh) > 0 {
		profit = (profit + heap.Pop(&mh).(int)) % mod
	}

	return profit
}

type carInfo struct {
	deadline int
	profit   int
}

type minIntHeap []int

func (h minIntHeap) Len() int           { return len(h) }
func (h minIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minIntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *minIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

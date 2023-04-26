package queue

/*
Problem Description
You are given a matrix A of size N x 2 which represents different operations.
Assume initially you have a stack-like data structure you have to perform operations on it.

Operations are of two types:

1 x: push an integer x onto the stack and return -1.

2 0: remove and return the most frequent element in the stack.

If there is a tie for the most frequent element, the element closest to the top of the stack is removed and returned.

A[i][0] describes the type of operation to be performed. A[i][1] describe the element x or 0 corresponding to the operation performed.

Problem Constraints
1 <= N <= 100000
1 <= A[i][0] <= 2
0 <= A[i][1] <= 109

Input Format
The only argument given is the integer array A.

Output Format
Return the array of integers denoting the answer to each operation.

Example Input
Input 1:
A = [

	[1, 5]
	[1, 7]
	[1, 5]
	[1, 7]
	[1, 4]
	[1, 5]
	[2, 0]
	[2, 0]
	[2, 0]
	[2, 0]  ]

Input 2:
A =  [

	[1, 5]
	[2, 0]
	[1, 4]   ]

Example Output
Output 1:
[-1, -1, -1, -1, -1, -1, 5, 7, 5, 4]

Output 2:
[-1, 5, -1]

Example Explanation
Explanation 1:
Just simulate given operations.

Explanation 2:
Just simulate given operations.
*/
func GetMaxFrequencyStack(A [][]int) []int {
	ans := make([]int, len(A))
	mfs := newMaxFreqStack()
	for i := range A {
		if A[i][0] == 1 {
			ans[i] = mfs.push(A[i][1])
		} else {
			ans[i] = mfs.pop()
		}
	}
	return ans
}

type maxFreqStack struct {
	maxFreq  int            //for storing the max frequency so far
	hashMap  map[int]int    //for storing freq against each element
	stackMap map[int]*stack //for storing the stack of elements against each frequency
}

func newMaxFreqStack() maxFreqStack {
	return maxFreqStack{
		hashMap:  make(map[int]int),
		stackMap: make(map[int]*stack),
	}
}

func (mfs *maxFreqStack) push(val int) int {
	mfs.hashMap[val]++
	freq := mfs.hashMap[val]
	//if freq is max so far
	if freq > mfs.maxFreq {
		mfs.maxFreq = freq
	}
	_, ok := mfs.stackMap[freq]
	//if the stack is absent, create it
	if !ok {
		mfs.stackMap[freq] = newStack()
	}
	//for the given frequency, update the stack with the current element
	mfs.stackMap[freq].push(val)
	return -1
}

func (mfs *maxFreqStack) pop() int {
	//retrieve the element stored at the top of the stack against the max frequency
	val := mfs.stackMap[mfs.maxFreq].topInt()
	mfs.stackMap[mfs.maxFreq].pop()
	mfs.hashMap[val]--
	//if the max frequency stack is empty, decrease the max frequency so far
	if mfs.stackMap[mfs.maxFreq].size() == 0 {
		mfs.maxFreq--
	}
	return val
}

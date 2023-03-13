package tries1

import (
	"fmt"
	"math"
	"sort"
)

/*
Problem Description
There is a dictionary A of N words, and ith word has a unique weight Wi.
Another string array B of size M contains the prefixes. For every prefix B[i], output atmost 5 words from the dictionary A that start with the same prefix.
Output the words in decreasing order of their weight.

NOTE: If there is no word that starts with the given prefix output -1.

Problem Constraints
1 <= T <= 5
1 <= N <= 20000
1 <= M <= 10000
1 <= Wi <= 106
1 <= length of any string either in A or B <= 30

Input Format
First line is an integer T denoting the number of test cases.
For each test case,

First line denotes two space seperated integer N and M.
Second line denotes N space seperated strings denoting the words in dictionary.
Third line denotes N space seperated integers denoting the weight of each word in the dictionary.
Fourth line denotes M space seperated integers denoting the prefixes.

Output Format
For every prefix B[i], print the space seperated output on a new line.
NOTE: After every output there should be a space.

Example Input
Input 1:
1
6 3
abcd aecd abaa abef acdcc acbcc
2 1 3 4 6 5
ab abc a

Input 2:
1
5 3
aaaa aacd abaa abaa aadcc
3 4 1 2 6
aa aba abc

Example Output
Output 1:
abef abaa abcd
abcd
acdcc acbcc abef abaa abcd

Output 2:
aadcc aacd aaaa
abaa abaa
-1

Example Explanation
Explanation 1:
For the prefix "ab" 3 words in the dictionary have same prefix: ("abcd" : 2, "abaa" : 3, "abef" : 4). Ouput them in decreasing order of weight.
For the prefix "abc" only 1 word in the dictionary have same prefix: ("abcd" : 2).
For the prefix "a" all 6 words in the dictionary have same prefix: ("abcd" : 2, "aecd" : 1, "abaa" : 3, "abef" : 4, "acdcc" : 6, "acbcc" : 5).
Since we can output atmost 5 words. Output top 5 weighted words in decreasing order of weight.

Explanation 2:
For the prefix "aa" 3 words in the dictionary have same prefix: ("aaaa" : 3, "aacd" : 4, "aadcc" : 6). Ouput them in decreasing order of weight.
For the prefix "aba" 2 words in the dictionary have same prefix: ("abaa" : 1, "abaa" : 2).
For the prefix "abc" there is no word in the dictionary which have same prefix. So print -1.
*/
func AutoComplete() {
	// YOUR CODE GOES HERE
	// Please take input and print output to standard input/output (stdin/stdout)
	// E.g. 'fmt.Scanf' for input & 'fmt.Printf' for output
	var testcases int
	var N, M int
	fmt.Scan(&testcases)
	for i := 0; i < testcases; i++ {
		fmt.Scan(&N, &M)
		words := make([]string, N)
		weights := make([]int, N)
		prefixes := make([]string, M)
		for j := 0; j < N; j++ {
			fmt.Scan(&words[j])
		}
		for j := 0; j < N; j++ {
			fmt.Scan(&weights[j])
		}
		for j := 0; j < M; j++ {
			fmt.Scan(&prefixes[j])
		}
		root := newWeightedTrie()
		for i := 0; i < N; i++ {
			root.insertWord(words[i], i, weights[i])
		}

		for i := 0; i < M; i++ {
			data := root.retrieveWord(prefixes[i], words)
			publish(data)
		}
	}

}

type weightedTrie struct {
	child  map[byte]*weightedTrie
	wInfos []wInfo
}

type wInfo struct {
	idx    int
	weight int
}

func newWeightedTrie() *weightedTrie {
	return &weightedTrie{
		child: make(map[byte]*weightedTrie),
	}
}

func (root *weightedTrie) insertWord(inp string, idx int, weight int) {
	temp := root
	for i := range inp {
		if temp.child[inp[i]] == nil {
			temp.child[inp[i]] = newWeightedTrie()
		}
		//insert the weight and the index
		temp.child[inp[i]].wInfos = append(temp.child[inp[i]].wInfos, wInfo{
			idx:    idx,
			weight: weight,
		})
		temp = temp.child[inp[i]]
	}
}

func (root *weightedTrie) retrieveWord(inp string, words []string) []string {
	var ans []string
	temp := root
	for i := range inp {
		if temp.child[inp[i]] == nil {
			return nil
		}
		temp = temp.child[inp[i]]
	}

	//sort the wighted Info array by weight in DESC order
	sort.Slice(temp.wInfos, func(i, j int) bool {
		return temp.wInfos[i].weight > temp.wInfos[j].weight
	})

	//take atmost 5
	for i := 0; i < min(len(temp.wInfos), 5); i++ {
		ans = append(ans, words[temp.wInfos[i].idx])
	}
	return ans
}

func min(args ...int) int {
	minVal := math.MaxInt
	for i := range args {
		if args[i] < minVal {
			minVal = args[i]
		}
	}
	return minVal
}

func publish(inp []string) {
	if len(inp) == 0 {
		fmt.Print(-1, " ")
		fmt.Println("")
		return
	}
	for i := range inp {
		fmt.Print(inp[i], " ")
	}
	fmt.Println("")
}

package tries1

/*
Problem Description
Given an array of words A (dictionary) and another array B (which contain some words).
You have to return the binary array (of length |B|) as the answer where 1 denotes that the word is present in the dictionary and 0 denotes it is not present.
Formally, for each word in B, you need to return 1 if it is present in Dictionary and 0 if not.
Such problems can be seen in real life when we work on any online editor (like Google Documnet), if the word is not valid it is underlined by a red line.

NOTE: Try to do this in O(n) time complexity.

Problem Constraints
1 <= |A| <= 1000
1 <= sum of all strings in A <= 50000
1 <= |B| <= 1000

Input Format
First argument is array of strings A.
First argument is array of strings B.

Output Format
Return the binary array of integers according to the given format.

Example Input
Input 1:
A = [ "hat", "cat", "rat" ]
B = [ "cat", "ball" ]

Input 2:
A = [ "tape", "bcci" ]
B = [ "table", "cci" ]

Example Output
Output 1:
[1, 0]

Output 2:
[0, 0]

Example Explanation
Explanation 1:
Only "cat" is present in the dictionary.

Explanation 2:
None of the words are present in the dictionary.
*/
func AreWordsPresent(A []string, B []string) []int {
	ans := make([]int, len(B))
	root := createDictionary(A)
	for i := range B {
		cur := root
		complete := true
		for j := range B[i] {
			_, ok := cur.child[B[i][j]]
			if !ok {
				complete = false
				break
			}
			cur = cur.child[B[i][j]]
		}
		if complete && cur.isEnd {
			ans[i]++
		}
	}
	return ans
}

func createDictionary(A []string) *trieNode {
	root := newTrieNode()
	for i := range A {
		cur := root
		for j := range A[i] {
			//if character is present at the cur node's map
			_, ok := cur.child[A[i][j]]
			if !ok {
				//if not, insert a new node against the given character
				cur.child[A[i][j]] = newTrieNode()
			}
			//move to the node against the given character
			cur = cur.child[A[i][j]]
		}
		cur.isEnd = true
	}
	return root
}

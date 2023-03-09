package tries1

import (
	"strings"
)

/*
Problem Description
Given two arrays of strings A of size N and B of size M.
Return a binary string C where C[i] = '1' if B[i] can be found in dictionary A using exactly one modification in B[i], Else C[i] = '0'.
NOTE: modification is defined as converting a character into another character.

Problem Constraints
1 <= N <= 30000
1 <= M <= 2500
1 <= length of any string either in A or B <= 20
strings contains only lowercase alphabets

Input Format
First argument is the string arrray A.
Second argument is the string array B.

Output Format
Return a binary string C where C[i] = '1' if B[i] can be found in dictionary A using one modification in B[i], Else C[i] = '0'.

Example Input
Input 1:
A = [data, circle, cricket]
B = [date, circel, crikket, data, circl]

Input 2:
A = [hello, world]
B = [hella, pello, pella]

Example Output
Output 1:
"10100"

Output 2:
"110"

Example Explanation
Explanation 1:
1. date = dat*(can be found in A)
2. circel =(cannot be found in A using exactly one modification)
3. crikket = cri*ket(can be found in A)
4. data = (cannot be found in A using exactly one modification)
5. circl = (cannot be found in A using exactly one modification)

Explanation 2:
Only pella cannot be found in A using only one modification.
*/
func GetPrefixExist(A []string, B []string) string {
	var ans strings.Builder
	root := getDict(A)
	for i := range B {
		if check(root, 0, B[i], false) {
			ans.WriteByte('1')
		} else {
			ans.WriteByte('0')
		}
	}
	return ans.String()
}

func getDict(inp []string) *trieNode {
	root := newTrieNode()
	for i := range inp {
		cur := root
		for j := range inp[i] {
			_, ok := cur.child[inp[i][j]]
			if !ok {
				cur.child[inp[i][j]] = newTrieNode()
			}
			cur = cur.child[inp[i][j]]
		}
		cur.isEnd = true
	}
	return root
}

func check(root *trieNode, beg int, word string, diffChar bool) bool {
	var ans bool
	// if word is longer than the tree
	if root == nil {
		return false
		//if the word is finished, return the flag
	} else if beg == len(word) {
		return diffChar
	}
	char := word[beg]

	//if the char is found, jump to next char
	if _, ok := root.child[char]; ok {
		ans = ans || check(root.child[char], beg+1, word, diffChar)
	}

	//if a different char has not been encountered yet, loop on in the map
	if !diffChar {
		for key := range root.child {
			if key != char {
				ans = ans || check(root.child[key], beg+1, word, true)
			}
		}
	}

	return ans
}

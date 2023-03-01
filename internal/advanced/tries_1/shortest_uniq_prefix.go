package tries1

import "strings"

func GetShortestUniqPrefix(A []string) []string {
	var ans []string
	var sb strings.Builder
	root := createCountDictionary(A)
	for i := range A {
		cur := root
		sb.Reset()
		for j := range A[i] {
			cur = cur.child[A[i][j]]
			sb.WriteByte(A[i][j])
			if cur.count == 1 {
				ans = append(ans, sb.String())
				break
			}
		}
	}
	return ans
}

func createCountDictionary(A []string) *trieNode {
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
			cur.count++
		}
	}
	return root
}

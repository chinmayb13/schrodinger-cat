package tries1

type trieNode struct {
	child map[byte]*trieNode
	isEnd bool
	count int
}

func newTrieNode() *trieNode {
	return &trieNode{
		child: make(map[byte]*trieNode),
	}
}

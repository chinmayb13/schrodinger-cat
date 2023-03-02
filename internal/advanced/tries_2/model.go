package tries2

import "math"

type btNode struct {
	value int
	left  *btNode
	right *btNode
	arr   []int
}

func newbtNode() *btNode {
	return &btNode{}
}

type info struct {
	xor int
	beg int
}

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
}

func newTreeNode(val int) *treeNode {
	return &treeNode{
		value: val,
	}
}

func getBits(inp []int) int {
	max := -1
	for i := range inp {
		if inp[i] > max {
			max = inp[i]
		}
	}
	if max == 0 {
		return 0
	}
	return int(math.Floor(math.Log2(float64(max)) + 1))
}

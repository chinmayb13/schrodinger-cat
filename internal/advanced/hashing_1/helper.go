package hashing1

type TreeMap struct {
	root *Node
}

type Node struct {
	key         int
	value       bool
	left, right *Node
}

func (tm *TreeMap) Put(key int) {
	tm.root = tm.put(tm.root, key)
}

func (tm *TreeMap) put(n *Node, key int) *Node {
	if n == nil {
		return &Node{key: key, value: true}
	}
	if n.key == key {
		n.value = true
	} else if n.key > key {
		n.left = tm.put(n.left, key)
	} else {
		n.right = tm.put(n.right, key)
	}
	return n
}

func (tm *TreeMap) LowerBound(key int) (int, bool) {
	n := tm.lowerBound(tm.root, key)
	if n == nil {
		return -1, false
	}
	return n.key, true
}

func (tm *TreeMap) lowerBound(n *Node, key int) *Node {
	if n == nil {
		return nil
	}
	if n.key == key {
		return n
	} else if n.key > key {
		node := tm.lowerBound(n.left, key)
		if node == nil {
			return n
		}
		return node
	} else {
		return tm.lowerBound(n.right, key)
	}
}

func (tm *TreeMap) LE(key int) (int, bool) {
	n := tm.le(tm.root, key)
	if n == nil {
		return -1, false
	}
	return n.key, true
}

func (tm *TreeMap) le(n *Node, key int) *Node {
	if n == nil {
		return nil
	}
	if n.key == key {
		return n
	} else if n.key > key {
		return tm.le(n.left, key)

	} else {
		node := tm.le(n.right, key)
		if node == nil {
			return n
		}
		return node
	}
}

func (tm *TreeMap) Delete(key int) {
	tm.root = tm.delete(tm.root, key)
}

func (tm *TreeMap) delete(root *Node, key int) *Node {
	if root == nil {
		return root
	}
	if key < root.key {
		root.left = tm.delete(root.left, key)
	} else if key > root.key {
		root.right = tm.delete(root.right, key)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}
		minNode := root.right
		for minNode.left != nil {
			minNode = minNode.left
		}
		root.key = minNode.key
		root.value = minNode.value
		root.right = tm.delete(root.right, minNode.key)
	}
	return root
}

func (tm *TreeMap) Get(key int) (interface{}, bool) {
	n := tm.get(tm.root, key)
	if n == nil {
		return nil, false
	}
	return n.value, true
}

func (tm *TreeMap) get(n *Node, key int) *Node {
	if n == nil {
		return nil
	}
	if n.key == key {
		return n
	} else if n.key > key {
		return tm.get(n.left, key)
	} else {
		return tm.get(n.right, key)
	}
}

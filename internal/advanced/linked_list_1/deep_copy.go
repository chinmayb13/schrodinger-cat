package linkedlist1

/*
Problem Description
A linked list A is given such that each node contains an additional random pointer which could point to any node in the list or NULL.
Return a deep copy of the list.

Problem Constraints
0 <= |A| <= 106

Input Format
The first argument of input contains a pointer to the head of linked list A.

Output Format
Return a pointer to the head of the required linked list.

Example Input
Given list
1 -> 2 -> 3
with random pointers going from
1 -> 3
2 -> 1
3 -> 1

Example Output
1 -> 2 -> 3
with random pointers going from
1 -> 3
2 -> 1
3 -> 1

Example Explanation
You should return a deep copy of the list. The returned answer should not contain the same node as the original list, but a copy of them. The pointers in the returned list should not link to any node in the original input list.
*/
type RandomListNode struct {
	value  int
	next   *RandomListNode
	random *RandomListNode
}

func CopyRandomList(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}
	insertCopy(head)
	updateRand(head)
	return updateNext(head)
}

func insertCopy(head *RandomListNode) {
	temp := head
	for temp != nil {
		newNode := newRandomListNode(temp.value)
		newNode.next = temp.next
		temp.next = newNode
		temp = newNode.next
	}
}

func updateRand(head *RandomListNode) {
	p1 := head
	p2 := head.next
	for p2.next != nil {
		if p1.random != nil {
			p2.random = p1.random.next
		}
		p1 = p2.next
		p2 = p1.next
	}
	if p1.random != nil {
		p2.random = p1.random.next
	}
}

func updateNext(head *RandomListNode) *RandomListNode {
	p1 := head
	p2 := head.next
	h2 := p2
	for p2.next != nil {
		p1.next = p2.next
		p2.next = p2.next.next
		p1 = p1.next
		p2 = p2.next
	}
	p1.next = nil
	return h2
}

func newRandomListNode(val int) *RandomListNode {
	node := new(RandomListNode)
	node.value = val
	return node
}

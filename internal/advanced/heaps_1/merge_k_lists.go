package heaps1

import "math"

/*
Problem Description
Given a list containing head pointers of N sorted linked lists.
Merge these given sorted linked lists and return them as one sorted list.

Problem Constraints
1 <= total number of elements in given linked lists <= 100000

Input Format
The first and only argument is a list containing N head pointers.

Output Format
Return a pointer to the head of the sorted linked list after merging all the given linked lists.

Example Input
Input 1:
1 -> 10 -> 20
4 -> 11 -> 13
3 -> 8 -> 9

Input 2:
10 -> 12
13
5 -> 6

Example Output
Output 1:
1 -> 3 -> 4 -> 8 -> 9 -> 10 -> 11 -> 13 -> 20

Output 2:
5 -> 6 -> 10 -> 12 ->13

Example Explanation
Explanation 1:
The resulting sorted Linked List formed after merging is 1 -> 3 -> 4 -> 8 -> 9 -> 10 -> 11 -> 13 -> 20.

Explanation 2:
The resulting sorted Linked List formed after merging is 5 -> 6 -> 10 -> 12 ->13.
*/
type ListNode struct {
	value int
	next  *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	//createMinHeap
	createMinListHeap(lists)
	//loop until all the elements are exhausted
	for lists[0] != nil {
		//keep adding minimum elements to tail
		cur.next = lists[0]
		cur = cur.next
		//move the min list to its next
		lists[0] = lists[0].next
		//heapify again
		minListHeapify(0, lists)
	}
	return dummyHead.next
}

func getLC(idx int, lists []*ListNode) *ListNode {
	lc := (idx << 1) + 1
	if lc >= len(lists) {
		return nil
	}
	return lists[lc]
}

func getRC(idx int, lists []*ListNode) *ListNode {
	rc := (idx << 1) + 2
	if rc >= len(lists) {
		return nil
	}
	return lists[rc]
}

func minListHeapify(idx int, lists []*ListNode) {
	if len(lists) < 2 {
		return
	}

	for {
		lc := getLC(idx, lists)
		rc := getRC(idx, lists)
		cur := lists[idx]
		//if cur, lc and rc are nil, getMinNode will return nil
		minNode := getMinNode(cur, lc, rc)
		if minNode == cur {
			return
		} else if minNode == lc {
			lists[idx], lists[(idx<<1)+1] = lc, cur
			idx = (idx << 1) + 1
		} else {
			lists[idx], lists[(idx<<1)+2] = rc, cur
			idx = (idx << 1) + 2
		}
	}
}

func getMinNode(nodes ...*ListNode) *ListNode {
	minNode := &ListNode{
		value: math.MaxInt,
	}
	for i := range nodes {
		if nodes[i] != nil && nodes[i].value < minNode.value {
			minNode = nodes[i]
		}
	}
	if minNode.value == math.MaxInt {
		return nil
	}
	return minNode
}

func createMinListHeap(lists []*ListNode) {
	for i := (len(lists) - 2) >> 1; i >= 0; i-- {
		minListHeapify(i, lists)
	}
}

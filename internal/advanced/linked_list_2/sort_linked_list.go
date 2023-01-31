package linkedlist2

/*
Problem Description
Sort a linked list, A in O(n log n) time using constant space complexity.

Problem Constraints
0 <= |A| = 105

Input Format
The first and the only arugment of input contains a pointer to the head of the linked list, A.

Output Format
Return a pointer to the head of the sorted linked list.

Example Input
Input 1:
A = [3, 4, 2, 8]

Input 2:
A = [1]

Example Output
Output 1:
[2, 3, 4, 8]

Output 2:
[1]

Example Explanation
Explanation 1:
The sorted form of [3, 4, 2, 8] is [2, 3, 4, 8].

Explanation 2:
The sorted form of [1] is [1].
*/
func SortList(A *listNode) *listNode {
	h := A
	if h == nil || h.next == nil {
		return h
	}
	mid := getMiddle(h)
	h1 := mid.next
	mid.next = nil
	h = SortList(h)
	h1 = SortList(h1)
	return mergeSortedLists(h, h1)

}

func mergeSortedLists(h1 *listNode, h2 *listNode) *listNode {
	dummyHead := listNode_new(-1)
	tail := dummyHead
	p1, p2 := h1, h2
	for p1 != nil && p2 != nil {
		if p1.value < p2.value {
			tail.next = p1
			p1 = p1.next
		} else {
			tail.next = p2
			p2 = p2.next
		}
		tail = tail.next
	}
	if p1 != nil {
		tail.next = p1
	} else {
		tail.next = p2
	}
	return dummyHead.next
}

func getMiddle(h *listNode) *listNode {
	s, f := h, h
	for f.next != nil && f.next.next != nil {
		s = s.next
		f = f.next.next
	}
	return s
}

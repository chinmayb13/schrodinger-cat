package linkedlist2

/*
Problem Description
Given a singly linked list A
A: A0 → A1 → … → An-1 → An
reorder it to:
A0 → An → A1 → An-1 → A2 → An-2 → …
You must do this in-place without altering the nodes' values.

Problem Constraints
1 <= |A| <= 106

Input Format
The first and the only argument of input contains a pointer to the head of the linked list A.

Output Format
Return a pointer to the head of the modified linked list.

Example Input
Input 1:
A = [1, 2, 3, 4, 5]

Input 2:
A = [1, 2, 3, 4]

Example Output
Output 1:
[1, 5, 2, 4, 3]

Output 2:
[1, 4, 2, 3]

Example Explanation
Explanation 1:
The array will be arranged to [A0, An, A1, An-1, A2].

Explanation 2:
The array will be arranged to [A0, An, A1, An-1, A2].
*/
func ReorderList(A *listNode) *listNode {
	h := A
	if h == nil || h.next == nil || h.next.next == nil {
		return h
	}
	mid := getMiddleElement(h)
	h1 := mid.next
	mid.next = nil
	h1 = reverse(h1)
	return mergeLists(h, h1)
}

func mergeLists(h1, h2 *listNode) *listNode {
	dummyHead := listNode_new(-1)
	cur := dummyHead
	p1, p2 := h1, h2
	var c1, c2 int
	for p1 != nil && p2 != nil {
		if c1 <= c2 {
			cur.next = p1
			p1 = p1.next
			c1++
		} else {
			cur.next = p2
			p2 = p2.next
			c2++
		}
		cur = cur.next
	}
	if p1 != nil {
		cur.next = p1
	} else {
		cur.next = p2
	}
	return dummyHead.next
}

func getMiddleElement(A *listNode) *listNode {
	s, f := A, A
	for f.next != nil && f.next.next != nil {
		s = s.next
		f = f.next.next
	}
	return s
}

func reverse(h *listNode) *listNode {
	if h == nil || h.next == nil {
		return h
	}
	var prev *listNode
	cur := h
	for cur != nil {
		nxt := cur.next
		cur.next = prev
		prev = cur
		cur = nxt
	}
	return prev
}

package linkedlist3

/*
Problem Description
Given a singly linked list A, determine if it's a palindrome. Return 1 or 0, denoting if it's a palindrome or not, respectively.

Problem Constraints
1 <= |A| <= 105

Input Format
The first and the only argument of input contains a pointer to the head of the given linked list.

Output Format
Return 0, if the linked list is not a palindrome.
Return 1, if the linked list is a palindrome.

Example Input
Input 1:
A = [1, 2, 2, 1]

Input 2:
A = [1, 3, 2]

Example Output
Output 1:
1

Output 2:
0

Example Explanation
Explanation 1:
The first linked list is a palindrome as [1, 2, 2, 1] is equal to its reversed form.

Explanation 2:
The second linked list is not a palindrom as [1, 3, 2] is not equal to [2, 3, 1].
*/
func IsPalindrome(A *listNode) int {
	h := A
	mid := getMiddle(h)
	h1 := mid.next
	mid.next = nil
	h1 = reverse(h1)
	return compare(h, h1)
}

func getMiddle(h *listNode) *listNode {
	s, f := h, h
	for f.next != nil && f.next.next != nil {
		s = s.next
		f = f.next.next
	}
	return s
}

func compare(h1, h2 *listNode) int {
	p1, p2 := h1, h2
	for p1 != nil && p2 != nil {
		if p1.value != p2.value {
			return 0
		}
		p1 = p1.next
		p2 = p2.next
	}
	return 1
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

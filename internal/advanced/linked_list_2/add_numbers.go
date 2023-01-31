package linkedlist2

/*
Problem Description
You are given two linked lists, A and B, representing two non-negative numbers.

The digits are stored in reverse order, and each of their nodes contains a single digit.

Add the two numbers and return it as a linked list.

Problem Constraints
1 <= |A|, |B| <= 105

Input Format
The first argument of input contains a pointer to the head of linked list A.
The second argument of input contains a pointer to the head of linked list B.

Output Format
Return a pointer to the head of the required linked list.

Example Input
Input 1:
A = [2, 4, 3]
B = [5, 6, 4]

Input 2:
A = [9, 9]
B = [1]

Example Output
Output 1:
[7, 0, 8]

Output 2:
[0, 0, 1]

Example Explanation
Explanation 1:
A = 342 and B = 465. A + B = 807.

Explanation 2:
A = 99 and B = 1. A + B = 100.
*/
func AddTwoNumbers(A *listNode, B *listNode) *listNode {
	p1, p2 := A, B
	var h, cur, node *listNode
	carry := 0
	sum := 0
	for p1 != nil || p2 != nil || carry > 0 {
		sum = carry
		if p1 != nil {
			sum += p1.value
			p1 = p1.next
		}
		if p2 != nil {
			sum += p2.value
			p2 = p2.next
		}
		node = listNode_new(sum % 10)
		if h == nil {
			h = node
			cur = h
		} else {
			cur.next = node
			cur = cur.next
		}
		carry = sum / 10
	}
	return h
}

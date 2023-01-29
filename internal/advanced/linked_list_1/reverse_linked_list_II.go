package linkedlist1

/*
Problem Description
Reverse a linked list A from position B to C.
NOTE: Do it in-place and in one-pass.

Problem Constraints
1 <= |A| <= 106
1 <= B <= C <= |A|

Input Format
The first argument contains a pointer to the head of the given linked list, A.
The second arugment contains an integer, B.
The third argument contains an integer C.

Output Format
Return a pointer to the head of the modified linked list.

Example Input
Input 1:
A = 1 -> 2 -> 3 -> 4 -> 5
B = 2
C = 4

Input 2:
A = 1 -> 2 -> 3 -> 4 -> 5
B = 1
C = 5

Example Output
Output 1:
1 -> 4 -> 3 -> 2 -> 5

Output 2:
5 -> 4 -> 3 -> 2 -> 1

Example Explanation
Explanation 1:
In the first example, we want to reverse the highlighted part of the given linked list : 1 -> 2 -> 3 -> 4 -> 5
Thus, the output is 1 -> 4 -> 3 -> 2 -> 5

Explanation 2:
In the second example, we want to reverse the highlighted part of the given linked list : 1 -> 4 -> 3 -> 2 -> 5
Thus, the output is 5 -> 4 -> 3 -> 2 -> 1
*/
func ReverseBetween(A *listNode, B int, C int) *listNode {
	if B == C {
		return A
	}
	var pivot, head, localHead, localTail *listNode
	var prev *listNode
	count := 1
	head = A
	cur := head
	for cur != nil {
		switch {
		case count == B-1:
			pivot = cur
		case count == B:
			localTail = cur
		case count > B && count <= C:
			for count <= C {
				if count == C {
					localHead = cur
				}
				nxt := cur.next
				cur.next = prev
				prev = cur
				cur = nxt
				count++
			}
			if pivot != nil {
				pivot.next = localHead
			} else {
				head = localHead
			}
			localTail.next = cur
		}
		if cur != nil {
			prev = cur
			cur = cur.next
		}
		count++
	}
	return head
}

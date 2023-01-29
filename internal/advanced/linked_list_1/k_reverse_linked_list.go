package linkedlist1

/*
Problem Description
Given a singly linked list A and an integer B, reverse the nodes of the list B at a time and return the modified linked list.

Problem Constraints
1 <= |A| <= 103
B always divides A

Input Format
The first argument of input contains a pointer to the head of the linked list.
The second arugment of input contains the integer, B.

Output Format
Return a pointer to the head of the modified linked list.

Example Input
Input 1:
A = [1, 2, 3, 4, 5, 6]
B = 2

Input 2:
A = [1, 2, 3, 4, 5, 6]
B = 3

Example Output
Output 1:
[2, 1, 4, 3, 6, 5]

Output 2:
[3, 2, 1, 6, 5, 4]

Example Explanation
Explanation 1:
For the first example, the list can be reversed in groups of 2.
[[1, 2], [3, 4], [5, 6]]
After reversing the K-linked list
[[2, 1], [4, 3], [6, 5]]

Explanation 2:
For the second example, the list can be reversed in groups of 3.
[[1, 2, 3], [4, 5, 6]]
After reversing the K-linked list
[[3, 2, 1], [6, 5, 4]]
*/
func KReverseList(A *listNode, B int) *listNode {
	if B == 1 {
		return A
	}
	var head, tail, prev *listNode
	var cur, prevTail *listNode
	count := 1
	cur = A
	for cur != nil {
		//reached the head of the given partition of B elements
		if count%B == 0 {
			//if it is the first partition, update head
			if count == B {
				head = cur
				//else update the tail to point to this head
			} else {
				tail.next = cur
			}

		}
		//reached the tail of the given partition of B elements
		//do not make the current element point to the prev element
		if count%B == 1 {
			tail = prevTail
			prevTail = cur
			prev = cur
			cur = cur.next
			count++
			continue
		}
		//make the cur point to prev
		nxt := cur.next
		cur.next = prev
		prev = cur
		cur = nxt
		count++
	}
	prevTail.next = nil
	return head
}

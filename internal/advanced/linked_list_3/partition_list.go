package linkedlist3

/*
Problem Description
Given a linked list A and a value B, partition it such that all nodes less than B come before nodes greater than or equal to B.
You should preserve the original relative order of the nodes in each of the two partitions.

Problem Constraints
1 <= |A| <= 106
1 <= A[i], B <= 109

Input Format
The first argument of input contains a pointer to the head to the given linked list.
The second argument of input contains an integer, B.

Output Format
Return a pointer to the head of the modified linked list.

Example Input
Input 1:
A = [1, 4, 3, 2, 5, 2]
B = 3

Input 2:
A = [1, 2, 3, 1, 3]
B = 2

Example Output
Output 1:
[1, 2, 2, 4, 3, 5]

Output 2:
[1, 1, 2, 3, 3]

Example Explanation
Explanation 1:
[1, 2, 2] are less than B wheread [4, 3, 5] are greater than or equal to B.

Explanation 2:
[1, 1] are less than B wheread [2, 3, 3] are greater than or equal to B.
*/
func Partition(A *listNode, B int) *listNode {
	if A == nil || A.next == nil {
		return A
	}
	cur := A
	//greaterNodes for nodes >=B
	greaterNodes := listNode_new(-1)
	gTail := greaterNodes
	//filterNodes for nodes<B
	filterNodes := listNode_new(-1)
	fTail := filterNodes
	for cur != nil {
		if cur.value >= B {
			gTail.next = cur
			cur = cur.next
			gTail = gTail.next
			gTail.next = nil
		} else {
			fTail.next = cur
			cur = cur.next
			fTail = fTail.next
		}
	}
	fTail.next = greaterNodes.next
	return filterNodes.next

}
func PartitionAlt(A *listNode, B int) *listNode {
	if A == nil || A.next == nil {
		return A
	}
	cur, h := A, A
	var prev *listNode
	target := B
	tail, size := getSizeAndTail(h)
	for i := 1; i <= size; i++ {
		//if value >= B
		if cur.value >= target && cur.next != nil {
			//if there is atleast one value before current element which is < B
			if prev != nil {
				//remove the linking
				prev.next = cur.next
			} else {
				//if not, make the next element of the current as head
				h = cur.next
			}
			//move the current element to the tail
			node := cur
			cur = cur.next
			tail.next = node
			tail = tail.next
			tail.next = nil
		} else {
			prev = cur
			cur = cur.next
		}

	}

	return h

}

func getSizeAndTail(h *listNode) (*listNode, int) {
	if h == nil || h.next == nil {
		if h == nil {
			return h, 0
		}
		return h, 1
	}
	tail := h
	size := 1
	for tail.next != nil {
		tail = tail.next
		size++
	}
	return tail, size
}

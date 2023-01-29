package linkedlist1

import "fmt"

/*
Design and implement a Linked List data structure.
A node in a linked list should have the following attributes - an integer value and a pointer to the next node. It should support the following operations:

insert_node(position, value) - To insert the input value at the given position in the linked list.
delete_node(position) - Delete the value at the given position from the linked list.
print_ll() - Print the entire linked list, such that each element is followed by a single space (no trailing spaces).
Note:

If an input position does not satisfy the constraint, no action is required.
Each print query has to be executed in a new line.

Problem Constraints
1 <= value <= 105
1 <= position <= n where, n is the size of the linked-list.

Input Format
First line contains an integer denoting number of cases, let's say t.
Next t line denotes the cases.

Output Format
When there is a case of print_ll(),  Print the entire linked list, such that each element is followed by a single space. There should not be any trailing space.
NOTE: You don't need to return anything.

Example Input
5
i 1 23
i 2 24
p
d 1
p

Example Output
23 24
24

Example Explanation
After first two cases linked list contains two elements 23 and 24.
At third case print: 23 24.
At fourth case delete value at first position, only one element left 24.
At fifth case print: 24.
*/
var head *listNode

func Insert_node(position, value int) {
	// @params position, integer
	// @params value, integer
	node := listNode_new(value)
	if position == 1 {
		node.next = head
		head = node
		return
	}
	temp := getPosNode(position - 1)
	if temp != nil {
		node.next = temp.next
		temp.next = node
	}
}

func Delete_node(position int) {
	// @params position, integer
	var temp *listNode
	if head == nil {
		return
	}
	if position == 1 {
		temp = head.next
		head.next = nil
		head = temp
		return
	}
	temp = getPosNode(position - 1)
	if temp != nil && temp.next != nil {
		delNode := temp.next
		temp.next = temp.next.next
		delNode.next = nil
	}
}

func Print_ll() {
	// Output each element followed by a space
	temp := head
	for temp.next != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println(temp.value)
}

func getPosNode(position int) *listNode {
	temp := head
	for i := 1; i < position && temp != nil; i++ {
		temp = temp.next
	}
	return temp
}

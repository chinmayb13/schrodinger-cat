package trees5

/*
Problem Description
Given a binary search tree represented by root A, write a function to find the Bth smallest element in the tree.

Problem Constraints
1 <= Number of nodes in binary tree <= 100000
0 <= node values <= 10^9

Input Format
First and only argument is head of the binary tree A.

Output Format
Return an integer, representing the Bth element.

Example Input
Input 1:
            2
          /   \
         1    3
B = 2

Input 2:
            3
           /
          2
         /
        1
B = 1

Example Output
Output 1:
2

Output 2:
1

Example Explanation
Explanation 1:
2nd element is 2.

Explanation 2:
1st element is 1.
*/
//idea: use Morris algo for inorder traversal to find Kth smallest
func Kthsmallest(A *treeNode, B int) int {
	var prev *treeNode
	count := 0
	cur := A
	for cur != nil {
		prev = nil
		temp := cur.left
		for temp != nil && temp != cur {
			prev = temp
			temp = temp.right
		}
		if prev != nil && prev.right == nil {
			prev.right = cur
			cur = cur.left
		} else if prev != nil && prev.right == cur {
			prev.right = nil
			count++
			if count == B {
				return cur.value
			}
			cur = cur.right
		} else {
			count++
			if count == B {
				return cur.value
			}
			cur = cur.right
		}
	}
	return -1
}

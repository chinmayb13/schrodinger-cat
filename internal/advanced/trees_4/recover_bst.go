package trees4

/*
Problem Description
Given a Binary Search Tree A. Also given are two nodes B and C. Find the lowest common ancestor of the two nodes.
Note 1 :- It is guaranteed that the nodes B and C exist.
Note 2 :- The LCA of B and C in A is the shared ancestor of B and C that is located farthest from the root.

Problem Constraints
1 <= Number of nodes in binary tree <= 105
1 <= B , C <= 105

Input Format
First argument is a root node of the binary tree, A.
Second argument is an integer B.
Third argument is an integer C.

Output Format
Return the LCA of the two nodes

Example Input
Input 1:

	       15
	     /    \
	   12      20
	   / \    /  \
	  10  14  16  27
	 /
	8

	B = 8
	C = 20

Input 2:

	       8
	      / \
	     6  21
	    / \
	   1   7

	B = 7
	C = 1

# Example Output

Output 1:
15

Output 2:
6

Example Explanation
Explanation 1:
The LCA of node 8 and 20 is the node 15.

Explanation 2:
The LCA of node 7 and 1 is the node 6.
*/
func RecoverTree(A *treeNode) []int {
	var ans []int
	curValue, lastValue := -1, -1
	cur := A
	var prev *treeNode
	for cur != nil {
		//pick the left subtree
		temp := cur.left
		prev = nil
		//go to most right of the left subtree
		for temp != nil && temp != cur {
			prev = temp
			temp = temp.right
		}
		//if the link is not created in leftsubtree, create the link and move the current to left
		if prev != nil && prev.right == nil {
			prev.right = cur
			cur = cur.left
			//if the link is already created, process the node, break the link and move the current to right
		} else if prev != nil && prev.right == cur {
			prev.right = nil
			lastValue = curValue
			curValue = cur.value
			//if there is a dip
			if lastValue > curValue {
				//if first dip, save both
				if len(ans) == 0 {
					ans = append(ans, curValue, lastValue)
					//if second dip, save the cur value
				} else {
					ans[0] = curValue
				}
			}
			cur = cur.right
			//no left subtree is present
		} else {
			lastValue = curValue
			curValue = cur.value
			//if there is a dip
			if lastValue > curValue {
				//if first dip, save both
				if len(ans) == 0 {
					ans = append(ans, curValue, lastValue)
					//if second dip, save the cur value
				} else {
					ans[0] = curValue
				}
			}
			cur = cur.right
		}
	}
	return ans

}

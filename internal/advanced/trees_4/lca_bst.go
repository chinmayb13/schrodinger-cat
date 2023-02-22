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

Example Output
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
func FindLCAinBST(A *treeNode, B int, C int) int {
	return findLCA(A, B, C)
}

func findLCA(root *treeNode, comp1, comp2 int) int {
	if root == nil {
		return -1
	}
	val := root.value
	if val == comp1 || val == comp2 {
		return val
	} else if val < comp1 && val < comp2 {
		return findLCA(root.right, comp1, comp2)
	} else if val > comp1 && val > comp2 {
		return findLCA(root.left, comp1, comp2)
	}
	return val
}

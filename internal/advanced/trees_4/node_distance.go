package trees4

/*
Problem Description
Given a binary search tree.
Return the distance between two nodes with given two keys B and C. It may be assumed that both keys exist in BST.

NOTE: Distance between two nodes is number of edges between them.

Problem Constraints
1 <= Number of nodes in binary tree <= 1000000
0 <= node values <= 109

Input Format
First argument is a root node of the binary tree, A.
Second argument is an integer B.
Third argument is an integer C.

Output Format
Return an integer denoting the distance between two nodes with given two keys B and C

Example Input
Input 1:
         5
       /   \
      2     8
     / \   / \
    1   4 6   11
 B = 2
 C = 11

Input 2:
         6
       /   \
      2     9
     / \   / \
    1   4 7   10
 B = 2
 C = 6

Example Output
Output 1:
3

Output 2:
1

Example Explanation
Explanation 1:
Path between 2 and 11 is: 2 -> 5 -> 8 -> 11. Distance will be 3.

Explanation 2:
Path between 2 and 6 is: 2 -> 6. Distance will be 1
*/
func GetNodeDistance(A *treeNode, B int, C int) int {
	if B == C {
		return 0
	}
	//find lowest common ancestor
	lca := findLCANode(A, B, C)
	//find distance of two nodes from LCA
	return findDistance(lca, B) + findDistance(lca, C)
}

func findLCANode(root *treeNode, comp1, comp2 int) *treeNode {
	if root == nil {
		return nil
	}
	val := root.value
	if val == comp1 || val == comp2 {
		return root
	} else if val < comp1 && val < comp2 {
		return findLCANode(root.right, comp1, comp2)
	} else if val > comp1 && val > comp2 {
		return findLCANode(root.left, comp1, comp2)
	} else {
		return root
	}
}

func findDistance(root *treeNode, target int) int {
	var path int
	if root == nil {
		return -1
	}
	val := root.value
	if val == target {
		return 0
	}
	if val < target {
		path = findDistance(root.right, target)
		if path == -1 {
			return -1
		}
		return 1 + path
	} else {
		path = findDistance(root.left, target)
		if path == -1 {
			return -1
		}
		return 1 + path
	}
}

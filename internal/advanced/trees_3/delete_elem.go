package trees3

/*
Problem Description
Given a Binary Search Tree(BST) A. If there is a node with value B present in the tree delete it and return the tree.
Note - If there are multiple options, always replace a node by its in-order predecessor

Problem Constraints
2 <= No. of nodes in BST <= 105
1 <= value of nodes <= 109
Each node has a unique value

Input Format
The first argument is the root node of a Binary Search Tree A.
The second argument is the value B.

Output Format
Delete the given node if found and return the root of the BST.

Example Input
Input 1:
            15
          /    \
        12      20
        / \    /  \
       10  14  16  27
      /
     8

     B = 10

Input 2:
            8
           / \
          6  21
         / \
        1   7

     B = 6

Example Output
Output 1:
            15
          /    \
        12      20
        / \    /  \
       8  14  16  27

Output 2:
            8
           / \
          1  21
           \
            7

Example Explanation
Explanation 1:
Node with value 10 is deleted

Explanation 2:
Node with value 6 is deleted
*/
func DeleteElem(A *treeNode, B int) *treeNode {
	return deleteElem(A, B)
}

func deleteElem(root *treeNode, elem int) *treeNode {
	if root == nil {
		return nil
	}
	if root.value < elem {
		root.right = deleteElem(root.right, elem)
	} else if root.value > elem {
		root.left = deleteElem(root.left, elem)
	} else {
		//if root has either 0 or 1 child
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
			//if root has 2 children
		} else {
			//look for the greatest element in left subtree and assign the value
			root.value = findGreatest(root.left)
			//delete the same element from the left subtree
			root.left = deleteElem(root.left, root.value)
		}
	}
	return root
}

func findGreatest(root *treeNode) int {
	temp := root
	for temp.right != nil {
		temp = temp.right
	}
	return temp.value
}

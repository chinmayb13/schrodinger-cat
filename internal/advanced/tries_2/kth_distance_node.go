package tries2

/*
Problem Description
Given the root of a binary tree A, the value of a target node B, and an integer C.
Return an array of the values of all nodes that have a distance C from the target node B.
Note :- The target node B always exist. All the nodes has unique value

Problem Constraints
1 <= number of nodes <= 105
1 <= B <= 105
1 <= C <= 105

Input Format
First argument is the root node of the binary tree, A.
Second argument is an integer denoting the value of the target node B.
Third argument is an integer denoting C.

Output Format
Return an integer array denoting the nodes at a distance C from targer node B

Example Input
Input 1:
A = 1

	 \
	  2
	 /
	3

B = 2
C = 1

Input 2:
A = 1

	 / \
	6   2
	   /
	  3

B = 6
C = 2

Example Output
Output 1:
[1, 3]

Output 2:
[2]

Example Explanation
Explanation 1:
The nodes 1 and 3 are at a distance 1 from node 2.

Explanation 2:
The node 2 is at a distance 2 from node 6.
*/
func GetKthNodes(A *treeNode, B int, C int) []int {
	var ans []int
	findNodes(A, C, B, &ans)
	return ans
}

func getNodesFromRoot(root *treeNode, k int, ans *[]int) {
	if root == nil {
		return
	}
	if k == 0 {
		*ans = append(*ans, root.value)
	}
	getNodesFromRoot(root.left, k-1, ans)
	getNodesFromRoot(root.right, k-1, ans)
}

func findNodes(root *treeNode, k, target int, ans *[]int) int {
	//if node or distance is invalid, return
	if root == nil || k < 0 {
		return 0
	}

	//if target is found, search for kth distant nodes in its subtree
	if root.value == target {
		getNodesFromRoot(root, k, ans)
		return 1
	}

	//search in left and right subtrees
	l := findNodes(root.left, k, target, ans)
	r := findNodes(root.right, k, target, ans)

	//if the current node is at kth distance, add it
	if l == k || r == k {
		*ans = append(*ans, root.value)
	}
	//if target is in left subtree, find kth distant nodes in right subtree
	if l > 0 {
		getNodesFromRoot(root.right, k-l-1, ans)
		return l + 1
		//if target is in right subtree, find kth distant nodes in left subtree
	} else if r > 0 {
		getNodesFromRoot(root.left, k-r-1, ans)
		return r + 1
	}
	return 0
}

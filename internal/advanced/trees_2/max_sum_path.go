package trees2

/*
Problem Description
Given a binary tree T, find the maximum path sum.

The path may start and end at any node in the tree.

Problem Constraints
1 <= Number of Nodes <= 7e4
-1000 <= Value of Node in T <= 1000

Input Format
The first and the only argument contains a pointer to the root of T, A.

Output Format
Return an integer representing the maximum sum path.

Example Input
Input 1:
    1
   / \
  2   3

Input 2:
       20
      /  \
   -10   20
        /  \
      -10  -50

Example Output
Output 1:
6

Output 2:
40

Example Explanation
Explanation 1:
The path with maximum sum is: 2 -> 1 -> 3

Explanation 2:
The path with maximum sum is: 20 -> 20
*/
func MaxPathSum(A *treeNode) int {
	pathSum := -1001
	_ = calcSum(A, &pathSum)
	return pathSum
}

func calcSum(root *treeNode, pathSum *int) int {
	if root == nil {
		return 0
	}
	value := root.value
	leftChild := max(0, calcSum(root.left, pathSum))
	rightChild := max(0, calcSum(root.right, pathSum))

	*pathSum = max(*pathSum, value+leftChild+rightChild)

	return value + max(leftChild, rightChild)
}

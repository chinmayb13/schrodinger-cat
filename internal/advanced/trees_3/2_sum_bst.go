package trees3

/*
Problem Description
Given a binary search tree A, where each node contains a positive integer, and an integer B, you have to find whether or not there exist two different nodes X and Y such that X.value + Y.value = B.
Return 1 to denote that two such nodes exist. Return 0, otherwise.

Problem Constraints
1 <= size of tree <= 100000
1 <= B <= 109

Input Format
First argument is the head of the tree A.
Second argument is the integer B.

Output Format
Return 1 if such a pair can be found, 0 otherwise.

Example Input
Input 1:
         10
         / \
        9   20
B = 19

Input 2:
          10
         / \
        9   20

B = 40

Example Output
Output 1:
1

Output 2:
0

Example Explanation
Explanation 1:
10 + 9 = 19. Hence 1 is returned.

Explanation 2:
No such pair exists.
*/
func IsSumExists(A *treeNode, B int) int {
	hs := make(map[int]bool)
	cs := checkSum(A, hs, B)
	if cs > 0 {
		return 1
	}
	return 0

}

func checkSum(root *treeNode, hs map[int]bool, target int) int {
	if root == nil {
		return 0
	}
	l := checkSum(root.left, hs, target)
	if l == 1 {
		return 1
	}
	r := checkSum(root.right, hs, target)
	if r == 1 || hs[target-root.value] {
		return 1
	}
	hs[root.value] = true
	return 0
}

package trees4

/*
Problem Description
Find the lowest common ancestor in an unordered binary tree A, given two values, B and C, in the tree.
Lowest common ancestor: the lowest common ancestor (LCA) of two nodes and w in a tree or directed acyclic graph (DAG) is the lowest (i.e., deepest) node that has both v and w as descendants.

Problem Constraints
1 <= size of tree <= 100000
1 <= B, C <= 109

Input Format
First argument is head of tree A.
Second argument is integer B.
Third argument is integer C.

Output Format
Return the LCA.

Example Input
Input 1:
      1
     /  \
    2    3
B = 2
C = 3

Input 2:
      1
     /  \
    2    3
   / \
  4   5
B = 4
C = 5

Example Output
Output 1:
1

Output 2:
2

Example Explanation
Explanation 1:
LCA is 1.

Explanation 2:
LCA is 2.
*/
func Lca(A *treeNode, B int, C int) int {
	ans := -1
	stB := newStack()
	stC := newStack()
	//find path of B and C
	findPath(A, stB, B)
	findPath(A, stC, C)
	//if either of them don't exist, return -1
	if stB.size() == 0 || stC.size() == 0 {
		return ans
	}

	//reverse the stack to start the path with the root
	revB := newStack()
	for stB.size() > 0 {
		revB.push(stB.topNode())
		stB.pop()
	}
	revC := newStack()
	for stC.size() > 0 {
		revC.push(stC.topNode())
		stC.pop()
	}

	//keep updating the answer till the anncestors are same
	for revB.size() > 0 && revC.size() > 0 && revB.topNode().value == revC.topNode().value {
		ans = revB.topNode().value
		revB.pop()
		revC.pop()
	}
	return ans
}

func findPath(root *treeNode, st *stack, target int) bool {
	if root == nil {
		return false
	}
	st.push(root)
	if root.value == target {
		return true
	}
	okL := findPath(root.left, st, target)
	okR := findPath(root.right, st, target)
	if okL || okR {
		return true
	}
	st.pop()
	return false
}

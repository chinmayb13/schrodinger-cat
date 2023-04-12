package graphs1

/*
Problem Description
Clone an undirected graph. Each node in the graph contains a label and a list of its neighbors.

Note: The test cases are generated in the following format (use the following format to use See Expected Output option):
First integer N is the number of nodes.
Then, N intgers follow denoting the label (or hash) of the N nodes.
Then, N2 integers following denoting the adjacency matrix of a graph, where Adj[i][j] = 1 denotes presence of an undirected edge between the ith and jth node, O otherwise.

Problem Constraints
1 <= Number of nodes <= 105

Input Format
First and only argument is a node A denoting the root of the undirected graph.

Output Format
Return the node denoting the root of the new clone graph.

Example Input
Input 1:
      1
    / | \
   3  2  4
        / \
       5   6

Input 2:
      1
     / \
    3   4
   /   /|\
  2   5 7 6

Example Output
Output 1:
Output will the same graph but with new pointers:
      1
    / | \
   3  2  4
        / \
       5   6

Output 2:
      1
     / \
    3   4
   /   /|\
  2   5 7 6

Example Explanation
Explanation 1:
We need to return the same graph, but the pointers to the node should be different.
*/
func CloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	//hashMap to check if node has already been created
	hashMap := make(map[int]*Node)
	//seed with first node
	hashMap[node.Val] = newNode(node.Val)
	q := newQueue()
	//enqueue first node
	q.enqueueBack(node)
	for q.size() > 0 {
		//pop from queue and get corresponding clone node
		front := q.frontNode()
		cloneFront := hashMap[front.Val]
		q.dequeueFront()
		//iterate over adjacency list
		for i := range front.Neighbors {
			label := front.Neighbors[i].Val
			//if node is not created, create one and push the actual node to queue
			if hashMap[label] == nil {
				hashMap[label] = newNode(label)
				q.enqueueBack(front.Neighbors[i])
			}
			//add the nodes to the clone front neighbours
			cloneFront.Neighbors = append(cloneFront.Neighbors, hashMap[label])
		}
	}
	return hashMap[node.Val]
}

func newNode(val int) *Node {
	return &Node{
		Val: val,
	}
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func (s *queue) frontNode() *Node {
	front := s.front()
	if front != nil {
		return front.(*Node)
	}
	return nil
}

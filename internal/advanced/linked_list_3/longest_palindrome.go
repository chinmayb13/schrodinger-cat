package linkedlist3

/*
Problem Description
Given a linked list of integers. Find and return the length of the longest palindrome list that exists in that linked list.
A palindrome list is a list that reads the same backward and forward.
Expected memory complexity : O(1)

Problem Constraints
1 <= length of the linked list <= 2000
1 <= Node value <= 100

Input Format
The only argument given is head pointer of the linked list.

Output Format
Return the length of the longest palindrome list.

Example Input
Input 1:
2 -> 3 -> 3 -> 3

Input 2:
2 -> 1 -> 2 -> 1 ->  2 -> 2 -> 1 -> 3 -> 2 -> 2

Example Output
Output 1:
3

Output 2:
5

Example Explanation
Explanation 1:
3 -> 3 -> 3 is largest palindromic sublist

Explanation 2:
2 -> 1 -> 2 -> 1 -> 2 is largest palindromic sublist.
*/
func GetLongestPalindrome(A *listNode) int {
	ans := 0
	h := A
	var prev *listNode
	cur := h
	for cur != nil {
		nxt := cur.next
		cur.next = prev
		ans = max(ans, palindromeCount(prev, nxt)+1)
		ans = max(ans, palindromeCount(cur, nxt))
		prev = cur
		cur = nxt
	}
	return ans
}

func palindromeCount(p1, p2 *listNode) int {
	count := 0
	for p1 != nil && p2 != nil {
		if p1.value != p2.value {
			break
		}
		count += 2
		p1 = p1.next
		p2 = p2.next
	}
	return count
}

func max(inp1, inp2 int) int {
	if inp1 > inp2 {
		return inp1
	}
	return inp2
}

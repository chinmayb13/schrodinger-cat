package graphs5

import (
	"container/list"
)

/*
Problem Description
Given any source point, (C, D) and destination point, (E, F) on a chess board of size A x B, we need to find whether Knight can move to the destination or not.
The above figure details the movements for a knight ( 8 possibilities ).
If yes, then what would be the minimum number of steps for the knight to move to the said point. If knight can not move from the source point to the destination point, then return -1.

NOTE: A knight cannot go out of the board.

Problem Constraints
1 <= A, B <= 500

Input Format
The first argument of input contains an integer A.
The second argument of input contains an integer B.
The third argument of input contains an integer C.
The fourth argument of input contains an integer D.
The fifth argument of input contains an integer E.
The sixth argument of input contains an integer F.

Output Format
If it is possible to reach the destination point, return the minimum number of moves.
Else return -1.

Example Input
Input 1:

	A = 8
	B = 8
	C = 1
	D = 1
	E = 8
	F = 8

Input 2:

	A = 2
	B = 4
	C = 2
	D = 1
	E = 4
	F = 4

Example Output
Output 1:
6

Output 2:
-1

Example Explanation
Explanation 1:
The size of the chessboard is 8x8, the knight is initially at (1, 1) and the knight wants to reach position (8, 8).
The minimum number of moves required for this is 6.

Explanation 2:
It is not possible to move knight to position (4, 4) from (2, 1)
*/
func Knight(A int, B int, C int, D int, E int, F int) int {
	N := A
	M := B
	matrix := make([][]int, N)
	for i := range matrix {
		matrix[i] = make([]int, M)
		for j := range matrix[i] {
			matrix[i][j] = -1
		}
	}

	//create direction array
	dir := [][]int{{-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, -1}, {2, 1}, {-1, -2}, {1, -2}}

	//distance of source to be 0
	matrix[C-1][D-1] = 0
	q := newQueue()

	//push source to the queue
	q.enqueueBack(&coord{row: C - 1, col: D - 1})
	for q.size() > 0 {
		front := q.frontCoord()
		q.dequeueFront()
		row := front.row
		col := front.col
		//iterate in all directions
		for i := 0; i < 8; i++ {
			newRow := row + dir[i][0]
			newCol := col + dir[i][1]

			//if valid unvisited cell, update path
			if newRow >= 0 && newRow < N && newCol >= 0 && newCol < M && matrix[newRow][newCol] == -1 {
				matrix[newRow][newCol] = matrix[row][col] + 1
				q.enqueueBack(&coord{row: newRow, col: newCol})
			}
		}
	}

	if E > N || F > M {
		return -1
	}

	//return distance to the destination
	return matrix[E-1][F-1]

}

type coord struct {
	row int
	col int
}
type queue struct {
	*list.List
}

func newQueue() *queue {
	return &queue{
		List: list.New(),
	}
}

func (s *queue) enqueueBack(inp interface{}) {
	s.PushBack(inp)
}

func (s *queue) dequeueFront() {
	if s.Len() > 0 {
		s.Remove(s.Front())
	}
}

func (s *queue) front() interface{} {
	if s.Len() > 0 {
		return s.Front().Value

	}
	return nil
}

func (s *queue) size() int {
	return s.Len()
}

func (s *queue) frontCoord() *coord {
	front := s.front()
	if front != nil {
		return front.(*coord)
	}
	return nil
}

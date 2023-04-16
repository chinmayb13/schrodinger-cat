package graphs2

import (
	"container/list"
	"math"
)

/*
Problem Description
Given a matrix of integers A of size N x M consisting of 0, 1 or 2.

Each cell can have three values:
The value 0 representing an empty cell.
The value 1 representing a fresh orange.
The value 2 representing a rotten orange.
Every minute, any fresh orange that is adjacent (Left, Right, Top, or Bottom) to a rotten orange becomes rotten. Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1 instead.

Note: Your solution will run on multiple test cases. If you are using global variables, make sure to clear them.

Problem Constraints
1 <= N, M <= 1000
0 <= A[i][j] <= 2

Input Format
The first argument given is the integer matrix A.

Output Format
Return the minimum number of minutes that must elapse until no cell has a fresh orange.

If this is impossible, return -1 instead.

Example Input
Input 1:
A = [   [2, 1, 1]
        [1, 1, 0]
        [0, 1, 1]   ]

Input 2:
A = [   [2, 1, 1]
        [0, 1, 1]
        [1, 0, 1]   ]

Example Output
Output 1:
4

Output 2:
-1

Example Explanation
Explanation 1:
Max of 4 using (0,0) , (0,1) , (1,1) , (1,2) , (2,2)

Explanation 2:
Task is impossible
*/

func GetMinimumTime(A [][]int) int {
	inp := A
	rows := len(inp)
	cols := len(inp[0])
	//create a time matrix
	timeMatrix := make([][]int, rows)
	for i := range timeMatrix {
		timeMatrix[i] = make([]int, cols)
		for j := range timeMatrix[i] {
			//prepopulate it with infinity
			timeMatrix[i][j] = math.MaxInt
		}
	}
	q := newQueue()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			//if rotten orange, enqueue it and set the time elapsed for it to be zero
			if inp[i][j] == 2 {
				q.enqueueBack(&coord{row: i, col: j})
				timeMatrix[i][j] = 0
			}
		}
	}

	//do BFS on 4 directions
	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for q.size() > 0 {
		front := q.frontCoord()
		q.dequeueFront()
		for k := 0; k < 4; k++ {
			newRow := front.row + dir[k][0]
			newCol := front.col + dir[k][1]
			//if the adjacent cell contains a fresh orange and the new time elapsed to make this orange rotten
			//is less than the previous time, enqueue the coordinate and set the new minimum time
			if isFresh(newRow, newCol, inp) && timeMatrix[front.row][front.col]+1 < timeMatrix[newRow][newCol] {
				q.enqueueBack(&coord{row: newRow, col: newCol})
				timeMatrix[newRow][newCol] = timeMatrix[front.row][front.col] + 1
			}
		}
	}

	timeElapsed := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if inp[i][j] == 1 {
				//if there is fresh orange which is unexplored, return -1
				if timeMatrix[i][j] == math.MaxInt {
					return -1
				}
				//else take the max time lapsed
				timeElapsed = max(timeElapsed, timeMatrix[i][j])
			}
		}
	}

	return timeElapsed
}

func isFresh(row, col int, inp [][]int) bool {
	rows := len(inp)
	cols := len(inp[0])
	return row >= 0 && row < rows && col >= 0 && col < cols && inp[row][col] == 1
}

func max(inp1, inp2 int) int {
	if inp1 > inp2 {
		return inp1
	}
	return inp2
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

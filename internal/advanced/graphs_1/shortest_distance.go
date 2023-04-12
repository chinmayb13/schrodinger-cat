package graphs1

import "math"

/*
Problem Description
Given a matrix of integers A of size N x M describing a maze. The maze consists of empty locations and walls.
1 represents a wall in a matrix and 0 represents an empty location in a wall.
There is a ball trapped in a maze. The ball can go through empty spaces by rolling up, down, left or right, but it won't stop rolling until hitting a wall (maze boundary is also considered as a wall). When the ball stops, it could choose the next direction.
Given two array of integers of size B and C of size 2 denoting the starting and destination position of the ball.
Find the shortest distance for the ball to stop at the destination. The distance is defined by the number of empty spaces traveled by the ball from the starting position (excluded) to the destination (included). If the ball cannot stop at the destination, return -1.

Problem Constraints
2 <= N, M <= 100
0 <= A[i] <= 1
0 <= B[i][0], C[i][0] < N
0 <= B[i][1], C[i][1] < M

Input Format
The first argument given is the integer matrix A.
The second argument given is an array of integer B.
The third argument if an array of integer C.

Output Format
Return a single integer, the minimum distance required to reach destination

Example Input
Input 1:
A = [ [0, 0], [0, 0] ]
B = [0, 0]
C = [0, 1]

Input 2:
A = [ [0, 0], [0, 1] ]
B = [0, 0]
C = [0, 1]

Example Output
Output 1:
1

Output 2:
1

Example Explanation
Explanation 1:
Go directly from start to destination in distance 1.

Explanation 2:
Go directly from start to destination in distance 1.
*/
func GetShortestPath(A [][]int, B []int, C []int) int {
	inp := A
	N := len(inp)
	M := len(inp[0])
	q := newQueue()
	distance := make([][]int, N)
	for i := 0; i < N; i++ {
		//initially set distance array with infinity
		distance[i] = make([]int, M)
		for j := range distance[i] {
			distance[i][j] = math.MaxInt
		}
	}
	//insert the starting point to the queue and set its distance as 0
	q.enqueueBack(coord{row: B[0], col: B[1]})
	distance[B[0]][B[1]] = 0

	//maintain a direction array to move
	rowDir := []int{-1, 0, 1, 0}
	colDir := []int{0, 1, 0, -1}
	for q.size() > 0 {
		front := q.frontCoord()
		q.dequeueFront()
		for k := 0; k < 4; k++ {
			count := 0
			newRow := front.row
			newCol := front.col
			//keep moving until a wall is encountered
			for checkCell(newRow+rowDir[k], newCol+colDir[k], N, M, inp) {
				newRow += rowDir[k]
				newCol += colDir[k]
				count++
			}

			//if distance calculated is minimum then the existing distance, update it
			if distance[front.row][front.col]+count < distance[newRow][newCol] {
				q.enqueueBack(coord{row: newRow, col: newCol})
				distance[newRow][newCol] = distance[front.row][front.col] + count
			}
		}
	}
	if distance[C[0]][C[1]] == math.MaxInt {
		return -1
	}
	return distance[C[0]][C[1]]
}

type coord struct {
	row int
	col int
}

func checkCell(row, col int, N, M int, inp [][]int) bool {
	return row >= 0 && row < N && col >= 0 && col < M && inp[row][col] == 0
}

func (s *queue) frontCoord() coord {
	front := s.front()
	if front != nil {
		return front.(coord)
	}
	return coord{}
}

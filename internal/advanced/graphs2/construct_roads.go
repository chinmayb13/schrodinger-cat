package graphs2

func GetRoadCount(A int, B [][]int) int {
	inp := A
	adjList := make([][]int, inp+1)
	colorArr := make([]int, inp+1)
	for i := range colorArr {
		colorArr[i] = -1
	}
	//create adj list
	for i := range B {
		adjList[B[i][0]] = append(adjList[B[i][0]], B[i][1])
		adjList[B[i][1]] = append(adjList[B[i][1]], B[i][0])
	}

	//do DFS
	searchRoadInDepth(1, 0, colorArr, adjList)
	set1, set2 := 0, 0
	for i := range colorArr {
		//count nodes of 0 color
		if colorArr[i] == 0 {
			set1++
			//count nodes of 1 color
		} else if colorArr[i] == 1 {
			set2++
		}
	}
	var mod int = 1e9 + 7
	//ans will be all possibilities - already present edges
	return (set1*set2)%mod - len(B)
}

func searchRoadInDepth(vertice, color int, colorArr []int, adjList [][]int) {
	//set color
	colorArr[vertice] = color
	for i := range adjList[vertice] {
		v := adjList[vertice][i]
		//if unvisited and do DFS
		if colorArr[v] == -1 {
			searchRoadInDepth(v, 1^color, colorArr, adjList)
		}
	}
}

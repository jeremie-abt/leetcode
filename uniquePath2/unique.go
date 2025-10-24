package uniquePath2

var m, n int

func recursion(obstacleGrid [][]int, lineIter, colIter int) int {
	// Condition de sortie négative, on ne peut pas passer par cette case.
	if obstacleGrid[lineIter][colIter] == 1 {
		return 0
	}

	// Condition de sortie positive : on a trouvé un chemin.
	if m-1 == lineIter && n-1 == colIter {
		return 1
	}

	var result int

	if lineIter < m-1 {
		result += recursion(obstacleGrid, lineIter+1, colIter)
	}
	if colIter < n-1 {
		result += recursion(obstacleGrid, lineIter, colIter+1)
	}
	return result
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m = len(obstacleGrid)
	n = len(obstacleGrid[0])

	return recursion(obstacleGrid, 0, 0)
}

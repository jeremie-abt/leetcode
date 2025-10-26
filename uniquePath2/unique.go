package uniquePath2

func solve(obstacleGrid [][]int, lineIter, colIter int) int {
	obstacleGrid[0][0] = 1

	for i := 0; i < colIter; i++ {
		for j := 0; j < lineIter; j++ {
			if j == 0 && i == 0 {
				continue
			}
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}

			var leftIdx, aboveIdx, currentVal int
			leftIdx = j - 1
			aboveIdx = i - 1

			if leftIdx >= 0 {
				currentVal += obstacleGrid[i][leftIdx]
			}
			if aboveIdx >= 0 {
				currentVal += obstacleGrid[aboveIdx][j]
			}
			obstacleGrid[i][j] = currentVal
		}
	}

	return obstacleGrid[colIter-1][lineIter-1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	return solve(obstacleGrid, len(obstacleGrid[0]), len(obstacleGrid))
}

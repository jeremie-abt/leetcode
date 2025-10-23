package validSodoku

import (
	"math"
	"slices"
)

func isValidLine(array []byte) bool {
	var currentValue byte

	slices.Sort(array)
	for _, i := range array {
		if i != '.' && i == currentValue {
			return false
		}

		currentValue = i
	}
	return true
}

func areValidSquares(squares map[int][]byte) bool {
	for i := 0; i < 9; i++ {
		if !isValidLine(squares[i]) {
			return false
		}
	}
	return false
}

func areValidLines(board [][]byte, length int) bool {
	squareNums := make(map[int][]byte)

	for i := 0; i < 9; i++ {
		lineNums := make([]byte, 0, length)
		columnNums := make([]byte, 0, length)
		for j := 0; j < 9; j++ {
			squareNumKey := math.Floor(float64(i)/3) + math.Floor(float64(j)/3)

			squareNums[int(squareNumKey)] = append(squareNums[int(squareNumKey)], board[i][j])
			lineNums = append(lineNums, board[i][j])
			columnNums = append(columnNums, board[j][i])
		}

		if !isValidLine(lineNums) {
			return false
		}
		if !isValidLine(columnNums) {
			return false
		}
	}

	if !areValidSquares(squareNums) {
		return false
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	return areValidLines(board, 9)
}

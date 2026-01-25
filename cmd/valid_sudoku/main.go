package main

import (
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	uniqueRowDigitsMap := make([]map[byte]struct{}, 9)
	uniqueColumnDigitsMap := make([]map[byte]struct{}, 9)
	uniqueSquareDigitsMap := make([]map[byte]struct{}, 9)

	for i := 0; i < 9; i++ {
		uniqueRowDigitsMap[i] = make(map[byte]struct{})
		uniqueColumnDigitsMap[i] = make(map[byte]struct{})
		uniqueSquareDigitsMap[i] = make(map[byte]struct{})
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			number := board[i][j]
			if number == '.' {
				continue
			}

			if _, foundInRow := uniqueRowDigitsMap[i][number]; foundInRow {
				return false
			}

			if _, foundInColumn := uniqueColumnDigitsMap[j][number]; foundInColumn {
				return false
			}

			squareIndex := (i/3)*3 + j/3

			if _, foundInSquare := uniqueSquareDigitsMap[squareIndex][number]; foundInSquare {
				return false
			}

			uniqueRowDigitsMap[i][number] = struct{}{}
			uniqueColumnDigitsMap[j][number] = struct{}{}
			uniqueSquareDigitsMap[squareIndex][number] = struct{}{}
		}
	}

	return true
}

func isValidSudokuLame(board [][]byte) bool {
	// Rows
	for _, row := range board {
		uniqueDigits := make(map[byte]struct{})
		for _, number := range row {
			if number == '.' {
				continue
			}
			if _, found := uniqueDigits[number]; found {
				return false
			}
			uniqueDigits[number] = struct{}{}
		}
	}

	// Cols
	for i := 0; i < 9; i++ {
		uniqueDigits := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			number := board[j][i]
			if number == '.' {
				continue
			}
			if _, found := uniqueDigits[number]; found {
				return false
			}
			uniqueDigits[number] = struct{}{}
		}
	}

	// Boxes
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			endI := i + 3
			endJ := j + 3
			uniqueDigits := make(map[byte]struct{})
			for subI := i; subI < endI; subI++ {
				for subJ := j; subJ < endJ; subJ++ {
					number := board[subI][subJ]
					if number == '.' {
						continue
					}
					if _, found := uniqueDigits[number]; found {
						return false
					}
					uniqueDigits[number] = struct{}{}
				}
			}
		}
	}

	return true
}

func main() {
	sudoku_one := [][]byte{{'1', '2', '.', '.', '3', '.', '.', '.', '.'},
		{'4', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '.', '3'},
		{'5', '.', '.', '.', '6', '.', '.', '.', '4'},
		{'.', '.', '.', '8', '.', '3', '.', '.', '5'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '.', '.', '.', '.', '.', '2', '.', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '8'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	fmt.Printf("sudoku_one should be valid, we got %v\n", isValidSudokuLame(sudoku_one))
	fmt.Printf("sudoku_one should be valid, we got %v\n", isValidSudoku(sudoku_one))
}

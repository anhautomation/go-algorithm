package algorithms

import (
	"strings"
)

func SolveNQueens(n int) ([][]string) {
	var solutions [][]string
	board := make([]string, n)
	for i := range board {
		board[i] = strings.Repeat(".", n)
	}
	backtrack(&solutions, board, 0, n)
	return solutions
}

func backtrack(solutions *[][]string, board []string, row, n int) {
	if row == n {
		*solutions = append(*solutions, append([]string(nil), board...))
		return
	}
	for col := 0; col < n; col++ {
		if isValid(board, row, col, n) {
			board[row] = board[row][:col] + "Q" + board[row][col+1:]
			backtrack(solutions, board, row+1, n)
			board[row] = board[row][:col] + "." + board[row][col+1:]
		}
	}
}

func isValid(board []string, row, col, n int) (bool) {
	// Check if placing a queen at board[row][col] is valid.
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
		if col-row+i >= 0 && board[i][col-row+i] == 'Q' {
			return false
		}
		if col+row-i < n && board[i][col+row-i] == 'Q' {
			return false
		}
	}
	return true
}
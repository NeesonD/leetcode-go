package _037_solveSudoku

func solveSudoku(board [][]byte) {
	help(board, 0, 0)
}

func help(board [][]byte, i int, j int) bool {
	m := 9
	n := 9
	if j == n {
		return help(board, i+1, 0)
	}
	if i == m {
		return true
	}
	if board[i][j] != '.' {
		return help(board, i, j+1)
	}
	for b := '1'; b <= '9'; b++ {
		if !isValid(board, i, j, byte(b)) {
			continue
		}
		board[i][j] = byte(b)
		if help(board, i, j+1) {
			return true
		}
		board[i][j] = '.'
	}
	return false
}

func isValid(board [][]byte, i int, j int, b byte) bool {
	for x := 0; x < 9; x++ {
		if board[i][x] == b {
			return false
		}
		if board[x][j] == b {
			return false
		}
		if board[(i/3)*3+x/3][(j/3)*3+x%3] == b {
			return false
		}
	}
	return true
}

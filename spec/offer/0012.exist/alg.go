package _012_exist

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[0] {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i, j, k int) bool {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[k] {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	board[i][j] = 1
	res := dfs(board, word, i+1, j, k+1) || dfs(board, word, i-1, j, k+1) ||
		dfs(board, word, i, j+1, k+1) ||
		dfs(board, word, i, j-1, k+1)
	board[i][j] = word[k]
	return res
}

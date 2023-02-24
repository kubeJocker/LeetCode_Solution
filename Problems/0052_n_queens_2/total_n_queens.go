package _051_n_queens

func totalNQueens(n int) int {
	res, col, row, diagonal1, diagonal2 := 0, make([]bool, n), []int{}, make([]bool, 2*n-1), make([]bool, 2*n-1)
	backtrack(n, &col, &row, &diagonal1, &diagonal2, &res)
	return res
}

func backtrack(n int, col *[]bool, row *[]int, diagonal1 *[]bool, diagonal2 *[]bool, res *int) {
	index := len(*row)
	if index == n {
		*res++
		return
	}
	for i := 0; i < n; i++ {
		if !(*col)[i] && !(*diagonal1)[index+i] && !(*diagonal2)[index-i+n-1] {
			*row = append(*row, i)
			(*col)[i], (*diagonal1)[index+i], (*diagonal2)[index-i+n-1] = true, true, true
			backtrack(n, col, row, diagonal1, diagonal2, res)
			*row = (*row)[:len(*row)-1]
			(*col)[i], (*diagonal1)[index+i], (*diagonal2)[index-i+n-1] = false, false, false
		}
	}
	return
}

func generateBoard(n int, row *[]int) []string {
	board := []string{}
	res := ""
	for i := 0; i < n; i++ {
		res += "."
	}
	for i := 0; i < n; i++ {
		board = append(board, res)
	}
	for i := 0; i < n; i++ {
		tmp := []byte(board[i])
		tmp[(*row)[i]] = 'Q'
		board[i] = string(tmp)
	}
	return board
}

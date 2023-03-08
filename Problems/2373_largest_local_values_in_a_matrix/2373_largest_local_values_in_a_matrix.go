package _373_largest_local_values_in_a_matrix

func largestLocal(grid [][]int) [][]int {
	length := len(grid)
	if length < 3 {
		return nil
	}
	res := make([][]int, length-2)
	for i := 0; i < length-2; i++ {
		res[i] = make([]int, length-2)
		for j := 0; j < length-2; j++ {
			for r := i; r < i+3; r++ {
				for c := j; c < j+3; c++ {
					if res[i][j] < grid[r][c] {
						res[i][j] = grid[r][c]
					}
				}
			}
		}
	}
	return res
}

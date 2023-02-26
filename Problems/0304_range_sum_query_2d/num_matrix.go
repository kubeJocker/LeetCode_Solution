package _304_range_sum_query_2d

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	temp := make([][]int, m+1)
	if m == 0 || n == 0 {
		return NumMatrix{matrix: temp}
	}
	for index, _ := range temp {
		temp[index] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			temp[i][j] = temp[i-1][j] + temp[i][j-1] + matrix[i-1][j-1] - temp[i-1][j-1]
		}
	}
	return NumMatrix{matrix: temp}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.matrix[row2+1][col2+1] - this.matrix[row1][col2+1] -
		this.matrix[row2+1][col1] + this.matrix[row1][col1]
}

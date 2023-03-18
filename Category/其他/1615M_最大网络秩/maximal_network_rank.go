package _615M_最大网络秩

// 设计到的边减共同的边
func maximalNetworkRank(n int, roads [][]int) int {
	degree, flags := make([]int, n), make([][]int, n)
	for i := range flags {
		flags[i] = make([]int, n)
	}
	for _, road := range roads {
		flags[road[0]][road[1]] = 1
		flags[road[1]][road[0]] = 1
		degree[road[0]]++
		degree[road[1]]++
	}
	ans := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			num := degree[i] + degree[j] - flags[i][j]
			if ans < num {
				ans = num
			}
		}
	}
	return ans
}

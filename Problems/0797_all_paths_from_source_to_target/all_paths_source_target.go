package _797_all_paths_from_source_to_target

func allPathsSourceTarget(graph [][]int) [][]int {
	stack, res := []int{0}, [][]int{}
	var dfs func(int2 int)
	dfs = func(x int) {
		if x == len(graph)-1 {
			res = append(res, append([]int(nil), stack...))
			return
		}
		for _, y := range graph[x] {
			stack = append(stack, y)
			dfs(y)
			stack = stack[:len(stack)-1]
		}
	}
	dfs(0)
	return res
}

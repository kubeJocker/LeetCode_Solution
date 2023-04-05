package _785_is_graph_bipartite

func isBipartite(graph [][]int) bool {
	ok := true
	// 记录图中节点的颜色，false 和 true 代表两种不同颜色
	color := make([]bool, len(graph))
	// 记录图中节点是否被访问过
	visited := make([]bool, len(graph))

	// 因为图不一定是联通的，可能存在多个子图
	// 所以要把每个节点都作为起点进行一次遍历
	// 如果发现任何一个子图不是二分图，整幅图都不算二分图
	var traverse func([][]int, int, []bool, []bool, *bool)
	traverse = func(graph [][]int, v int, visited []bool, color []bool, ok *bool) {
		if !*ok {
			return
		}
		visited[v] = true
		for _, w := range graph[v] {
			if !visited[w] {
				color[w] = !color[v]
				traverse(graph, w, visited, color, ok)
			} else {
				if color[w] == color[v] {
					*ok = false
					return
				}
			}
		}
	}
	for v := 0; v < len(graph); v++ {
		if !visited[v] {
			traverse(graph, v, visited, color, &ok)
		}
	}
	return ok
}

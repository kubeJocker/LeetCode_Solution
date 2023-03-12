package _617H_统计子树中城市之间最大距离

func countSubgraphsForEachDiameter(n int, edges [][]int) []int {

	//建立图
	graph := make([][]int, n)
	for _, edge := range edges {
		x, y := edge[0]-1, edge[1]-1
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}
	//求树的直径
	var inSet, vis [15]bool
	var diameter int
	var dfs func(int) int
	dfs = func(x int) (maxLen int) {
		vis[x] = true
		for _, y := range graph[x] {
			if !vis[y] && inSet[y] {
				ml := dfs(y) + 1
				diameter = max(diameter, maxLen+ml)
				maxLen = max(maxLen, ml)
			}
		}
		return
	}

	ans := make([]int, n-1)
	var f func(int)
	//枚举子树
	f = func(i int) {
		//选择城市完毕
		if i == n {
			//对于每个选中的城市，DFS测试连通性，计算树的直径
			for v, b := range inSet {
				if b {
					vis, diameter = [15]bool{}, 0
					dfs(v)
					break
				}
			}
			if diameter > 0 && vis == inSet {
				ans[diameter-1]++
			}
			return
		}
		//不选城市i
		f(i + 1)
		inSet[i] = true
		//选择城市i
		f(i + 1)
		//恢复
		inSet[i] = false
	}
	f(0)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

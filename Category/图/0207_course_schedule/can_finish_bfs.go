package _207_course_schedule

func canFinish_bfs(numCourses int, prerequisites [][]int) bool {
	// 建图，有向边代表「被依赖」关系
	graph := buildGraph(numCourses, prerequisites)
	// 构建入度数组
	indegree := make([]int, numCourses)
	for _, edge := range prerequisites {
		// 节点 to 的入度加一
		indegree[edge[0]]++
	}

	// 根据入度初始化队列中的节点
	q := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			// 节点 i 没有入度，即没有依赖的节点
			// 可以作为拓扑排序的起点，加入队列
			q = append(q, i)
		}
	}

	// 记录遍历的节点个数
	count := 0
	// 开始执行 BFS 循环
	for len(q) > 0 {
		// 弹出节点 cur，并将它指向的节点的入度减一
		cur := q[0]
		q = q[1:]
		count++
		for _, next := range graph[cur] {
			indegree[next]--
			if indegree[next] == 0 {
				// 如果入度变为 0，说明 next 依赖的节点都已被遍历
				q = append(q, next)
			}
		}
	}

	// 如果所有节点都被遍历过，说明不成环
	return count == numCourses
}

// 建图函数buildGraph

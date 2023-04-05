package _207_course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(numCourses, prerequisites)
	visited := make([]bool, numCourses)
	onPath := make([]bool, numCourses)
	hasCycle := false

	for i := 0; i < numCourses; i++ {
		// 遍历图中的所有节点
		traverse(graph, i, &visited, &onPath, &hasCycle)
	}
	return !hasCycle
}

func traverse(graph [][]int, s int, visited *[]bool, onPath *[]bool, hasCycle *bool) {
	if (*onPath)[s] {
		// 出现环
		*hasCycle = true
		return
	}
	if (*visited)[s] || *hasCycle {
		// 如果已经找到了环，也不用再遍历了
		return
	}
	(*visited)[s] = true
	(*onPath)[s] = true
	for _, t := range graph[s] {
		traverse(graph, t, visited, onPath, hasCycle)
	}
	(*onPath)[s] = false
}

func buildGraph(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = []int{}
	}
	for _, edge := range prerequisites {
		from, to := edge[1], edge[0]
		graph[from] = append(graph[from], to)
	}
	return graph
}

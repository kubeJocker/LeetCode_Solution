package _210_course_schedule_2

func findOrder(numCourses int, prerequisites [][]int) []int {
	postorder := []int{}
	graph := buildGraph(numCourses, prerequisites)
	visited := make([]bool, numCourses)
	onPath := make([]bool, numCourses)
	hasCycle := false

	for i := 0; i < numCourses; i++ {
		// 遍历图中的所有节点
		traverse(graph, i, &visited, &postorder, &onPath, &hasCycle)
	}
	if hasCycle {
		return []int{}
	}

	// 逆序后序遍历结果
	reverse(postorder)

	return postorder
}

func traverse(graph [][]int, s int, visited *[]bool, postorder *[]int, onPath *[]bool, hasCycle *bool) {
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
		traverse(graph, t, visited, postorder, onPath, hasCycle)
	}
	*postorder = append(*postorder, s)
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

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

package _514_path_with_maximum_probability

import "container/heap"

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	graph := make([][]tuple, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]tuple, 0)
	}
	// 构造邻接表结构表示图
	for i, edge := range edges {
		from := edge[0]
		to := edge[1]
		weight := succProb[i]
		// 无向图就是双向图
		graph[from] = append(graph[from], tuple{to, weight})
		graph[to] = append(graph[to], tuple{from, weight})
	}
	return dijkstra(start, end, graph)
}

type tuple struct {
	to   int
	prob float64
}

type state struct {
	// 图节点的 id
	id int
	// 从 start 节点到达当前节点的概率
	probFromStart float64
}

type pq []state

func (p pq) Len() int { return len(p) }

func (p pq) Less(i, j int) bool {
	return p[i].probFromStart > p[j].probFromStart
}

func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pq) Push(x interface{}) {
	item := x.(state)
	*p = append(*p, item)
}

func (p *pq) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	*p = old[0 : n-1]
	return item
}

func dijkstra(start, end int, graph [][]tuple) float64 {
	// 定义：probTo[i] 的值就是节点 start 到达节点 i 的最大概率
	probTo := make([]float64, len(graph))
	// base case，start 到 start 的概率就是 1
	probTo[start] = 1.0

	// 优先级队列，probFromStart 较大的排在前面
	pq := make(pq, 0)
	heap.Push(&pq, state{start, 1.0})
	for len(pq) > 0 {
		curState := heap.Pop(&pq).(state)
		curNodeID := curState.id
		curProbFromStart := curState.probFromStart

		// 遇到终点提前返回
		if curNodeID == end {
			return curProbFromStart
		}

		if curProbFromStart < probTo[curNodeID] {
			// 已经有一条概率更大的路径到达 curNode 节点了
			continue
		}

		// 将 curNode 的相邻节点装入队列
		for _, neighbor := range graph[curNodeID] {
			nextNodeID := neighbor.to
			// 看看从 curNode 达到 nextNode 的概率是否会更大
			probToNextNode := probTo[curNodeID] * neighbor.prob
			if probTo[nextNodeID] < probToNextNode {
				probTo[nextNodeID] = probToNextNode
				heap.Push(&pq, state{nextNodeID, probToNextNode})
			}
		}
	}
	// 如果到达这里，说明从 start 开始无法到达 end，返回 0
	return 0.0
}

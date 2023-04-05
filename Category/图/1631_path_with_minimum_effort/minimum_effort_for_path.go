package _631_path_with_minimum_effort

import (
	"container/heap"
	"math"
)

type State struct {
	x, y, effortFromStart int
}

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	// 定义：从 (0, 0) 到 (i, j) 的最小体力消耗是 effortTo[i][j]
	effortTo := make([][]int, m)
	for i := range effortTo {
		effortTo[i] = make([]int, n)
		for j := range effortTo[i] {
			effortTo[i][j] = math.MaxInt32
		}
	}
	// base case，起点到起点的最小消耗就是 0
	effortTo[0][0] = 0

	// 优先级队列，effortFromStart 较小的排在前面
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &State{0, 0, 0})

	// 从起点 (0, 0) 开始进行 BFS
	for pq.Len() > 0 {
		curState := heap.Pop(&pq).(*State)
		curX, curY, curEffortFromStart := curState.x, curState.y, curState.effortFromStart

		// 到达终点提前结束
		if curX == m-1 && curY == n-1 {
			return curEffortFromStart
		}

		if curEffortFromStart > effortTo[curX][curY] {
			continue
		}

		// 将 (curX, curY) 的相邻坐标装入队列
		for _, neighbor := range adj(heights, curX, curY) {
			nextX, nextY := neighbor[0], neighbor[1]
			// 计算从 (curX, curY) 达到 (nextX, nextY) 的消耗
			effortToNextNode := max(
				effortTo[curX][curY],
				abs(heights[curX][curY]-heights[nextX][nextY]),
			)
			// 更新 dp table
			if effortTo[nextX][nextY] > effortToNextNode {
				effortTo[nextX][nextY] = effortToNextNode
				heap.Push(&pq, &State{nextX, nextY, effortToNextNode})
			}
		}
	}
	// 正常情况不会达到这个 return
	return -1
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].effortFromStart < pq[j].effortFromStart
}

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*State)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func adj(heights [][]int, x, y int) [][]int {
	m, n := len(heights), len(heights[0])
	res := [][]int{}
	if x > 0 {
		res = append(res, []int{x - 1, y})
	}
	if y > 0 {
		res = append(res, []int{x, y - 1})
	}
	if x < m-1 {
		res = append(res, []int{x + 1, y})
	}
	if y < n-1 {
		res = append(res, []int{x, y + 1})
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

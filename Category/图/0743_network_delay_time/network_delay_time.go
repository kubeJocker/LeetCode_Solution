package _743_network_delay_time

import (
	"container/heap"
	"math"
)

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make([][]State, n)

	for _, edge := range times {
		from, to, weight := edge[0]-1, edge[1]-1, edge[2]
		graph[from] = append(graph[from], State{to, weight})
	}
	distTo := dijkstra(k-1, &graph)
	res := 0
	for i := 0; i < len(distTo); i++ {
		if distTo[i] == math.MaxInt {
			return -1
		}
		res = max(res, distTo[i])
	}
	return res
}

type State struct {
	id            int
	distFromStart int
}

type PriorityQueue []*State

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distFromStart < pq[j].distFromStart
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*State))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func dijkstra(start int, graph *[][]State) []int {
	distTo := make([]int, len(*graph))
	for i := range distTo {
		distTo[i] = math.MaxInt
	}
	distTo[start] = 0
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &State{id: start, distFromStart: 0})
	for len(pq) > 0 {
		curState := heap.Pop(&pq).(*State)
		curId := curState.id
		curDist := curState.distFromStart
		if curDist > distTo[curId] {
			continue
		}
		for _, neighbor := range (*graph)[curId] {
			nextNodeID := neighbor.id
			distToNextNode := distTo[curId] + neighbor.distFromStart
			// 更新 dp table
			if distTo[nextNodeID] > distToNextNode {
				distTo[nextNodeID] = distToNextNode
				heap.Push(&pq, &State{id: nextNodeID, distFromStart: distToNextNode})
			}
		}
	}
	return distTo
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

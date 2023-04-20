package _787_cheapest_flight_within_k_stops

import (
	"container/heap"
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := make([][]edge, n)
	for i := 0; i < n; i++ {
		graph[i] = []edge{}
	}
	for _, flight := range flights {
		from := flight[0]
		to := flight[1]
		price := flight[2]
		graph[from] = append(graph[from], edge{to, price})
	}
	k++
	return dijkstra(graph, src, k, dst)
}

type State struct {
	id             int
	costFromSrc    int
	nodeNumFromSrc int
}

type edge struct {
	to, price int
}

func dijkstra(graph [][]edge, src int, k int, dst int) int {
	//从起点src到达节点i的最短路径权重
	distTo := make([]int, len(graph))
	//从起点src到达节点i的最小权重路径
	nodeNumTo := make([]int, len(graph))
	for i := range distTo {
		distTo[i] = math.MaxInt32
		nodeNumTo[i] = math.MaxInt32
	}
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &State{src, 0, 0})

	for pq.Len() > 0 {
		curState := heap.Pop(&pq).(*State)
		curNodeId := curState.id
		costFromSrc := curState.costFromSrc
		curNodeNumFromSrc := curState.nodeNumFromSrc

		if curNodeId == dst {
			//找到最短路径
			return costFromSrc
		}
		if curNodeNumFromSrc == k {
			continue
		}

		for _, neighbor := range graph[curNodeId] {
			nextNodeId := neighbor.to
			costToNextNode := costFromSrc + neighbor.price
			nextNodeNumFromSrc := curNodeNumFromSrc + 1
			if distTo[nextNodeId] > costToNextNode {
				distTo[nextNodeId] = costToNextNode
				nodeNumTo[nextNodeId] = nextNodeNumFromSrc
			}
			//剪枝
			if costToNextNode > distTo[nextNodeId] &&
				nextNodeNumFromSrc > nodeNumTo[nextNodeId] {
				continue
			}
			heap.Push(&pq, &State{nextNodeId, costToNextNode, nextNodeNumFromSrc})
		}
	}
	return -1
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].costFromSrc < pq[j].costFromSrc
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

package _584_min_cost_to_connect_all_points

import (
	"container/heap"
	"math"
)

func minCostConnectPoints_prim(points [][]int) int {
	return prim(points, 0)
}

func prim(points [][]int, start int) int {
	// 结果
	var res int

	size := len(points)
	// 构建邻接矩阵
	var g = make([][]int, size)
	for i := 0; i < size; i++ {
		g[i] = make([]int, size)
	}
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			dist := int(math.Abs(float64(points[i][0])-float64(points[j][0]))) + int(math.Abs(float64(points[i][1])-float64(points[j][1])))
			g[i][j] = dist
			g[j][i] = dist
		}
	}

	// 维护一个小根堆，存储未处理节点的到新图的距离，堆顶是最小距离
	lowCost := &hp{}

	// 起始节点是 start，更新 lowCost 数组
	for i := 0; i < size; i++ {
		if i == start {
			continue
		}

		// 将剩余节点与起始节点之间的距离保存起来，用于寻找距离最小的节点
		heap.Push(lowCost, e{
			idx:  i,
			dist: g[i][start],
		})
	}

	// 循环 size - 1 次，因为只剩下 size - 1 个节点未处理
	for i := 0; i < size-1; i++ {
		// 寻找下一个最短距离的节点，并从堆中删除
		next := heap.Pop(lowCost).(e)
		res += next.dist
		nextIdx := next.idx

		// 更新 lowCost 中的值，并维护堆
		for i := 0; i < lowCost.Len(); i++ {
			if lowCost.data[i].dist > g[lowCost.data[i].idx][nextIdx] {
				lowCost.data[i].dist = g[lowCost.data[i].idx][nextIdx]
				heap.Fix(lowCost, i)
			}
		}
	}

	return res
}

type e struct {
	idx  int
	dist int
}

type hp struct {
	data []e
}

func (h hp) Len() int            { return len(h.data) }
func (h hp) Less(i, j int) bool  { return h.data[i].dist < h.data[j].dist }
func (h hp) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *hp) Push(x interface{}) { h.data = append(h.data, x.(e)) }
func (h *hp) Pop() interface{} {
	size := h.Len()
	data := h.data[size-1]
	h.data = h.data[:size-1]
	return data
}

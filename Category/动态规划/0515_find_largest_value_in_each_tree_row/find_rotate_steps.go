package _515_find_largest_value_in_each_tree_row

import "math"

func findRotateSteps(ring string, key string) int {
	indexMap := make([][]int, 26) // ring中的字符在ring中的索引集合
	for i, c := range ring { // 遍历ring
		cInt := c - 'a'      // 用一个int数代表当前字符
		indexMap[cInt] = append(indexMap[cInt], i) // 索引i加入对应的索引集合
	}

	memo := make([][]int, len(ring)) // memo二维数组，长度为ring的长度
	for i, _ := range memo {
		memo[i] = make([]int, len(key)) // 子数组是数组，长度为key长度
		for j, _ := range memo[i] {
			memo[i][j] = -1	// 初始填充-1，区别于距离（非负）
		}
	}
	var dfs func(ringI, keyI int) int
	// 当前12点对齐了ringI，对齐key[keyI]到key末尾字符，所需的最少步数
	dfs = func(ringI, keyI int) int {
		if keyI == len(key) { // 扫描key字符串的指针越界，返回步数0
			return 0
		}
		if memo[ringI][keyI] != -1 { // memo中有直接用memo的
			return memo[ringI][keyI]
		}
		cur := key[keyI]    // 当前想对齐的key字符
		curInt := cur - 'a' // 对应的int数
		res := math.MaxInt32 // 当前子问题的解
		// 遍历key[keyI]在ring中的索引数组，targetI是当前要对齐的目标索引
		for _, targetI := range indexMap[curInt] {
			d1 := abs(ringI - targetI)  // 从对齐ringI到对齐targetI的距离1
			d2 := len(ring) - d1   // 从对齐ringI到对齐targetI的距离2
			curMin := min(d1, d2)  // 取二者中的较小值
			res = min(res, curMin + dfs(targetI, keyI+1))// 当前targetI的解，试图刷新res
		}
		memo[ringI][keyI] = res  // 算出一个子问题的解，存入memo
		return res
	}

	return len(key) + dfs(0, 0) // 递归的入口
}

func abs(a int) int {
	if a < 0 {return -a}
	return a
}
func min(a, b int) int {
	if a > b {return b}
	return a
}

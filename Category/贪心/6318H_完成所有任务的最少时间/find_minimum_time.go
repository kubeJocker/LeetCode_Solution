package _318H_完成所有任务的最少时间

import "sort"

/*
贪心算法，将区间右端排序，然后遍历区间
如果时段已被分配，则duration--
有未分配的时段，则尽量分配在区间末端
*/
func findMinimumTime(tasks [][]int) (ans int) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})
	run := [2001]bool{}
	for _, task := range tasks {
		start, end, dur := task[0], task[1], task[2]
		//统计已经在运行的区间
		for _, b := range run[start : end+1] {
			if b {
				dur--
			}
		}
		for i := end; dur > 0; i-- {
			if !run[i] {
				run[i] = true
				dur--
				ans++
			}
		}
	}
	return
}

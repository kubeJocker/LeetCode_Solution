package _626_无矛盾的最佳球队

import "sort"

func bestTeamScore(scores []int, ages []int) int {
	n := len(scores)
	people := make([][]int, n)
	for i := range scores {
		people[i] = []int{scores[i], ages[i]}
	}
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] < people[j][0] {
			return true
		} else if people[i][0] > people[j][0] {
			return false
		}
		return people[i][1] < people[j][1]
	})
	dp := make([]int, n)
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if people[j][1] <= people[i][1] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] += people[i][0]
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package _653_minimum_deletions_to_make_string_balanced

/*
考虑 sss 的最后一个字母：
1.如果它是 ‘b’，则无需删除，问题规模缩小，变成「使 s 的前 n−1 个字母平衡的最少删除次数」。
2.如果它是 ‘a’：
	删除它，则答案为「使 s 的前 n−1 个字母平衡的最少删除次数」加上 1。
	保留它，那么前面的所有 ‘b’ 都要删除；
状态转移:
如果第 i个字母是 ‘b’，则 f[i]=f[i−1]
如果第 i个字母是 ‘a’，则 f[i]=min⁡(f[i−1]+1,cntB)
*/

func minimumDeletions_dp(s string) int {
	dp, cntB := 0, 0
	for _, c := range s {
		if c == 'b' { // f 值不变
			cntB++
		} else {
			dp = min(dp+1, cntB)
		}
	}
	return dp
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

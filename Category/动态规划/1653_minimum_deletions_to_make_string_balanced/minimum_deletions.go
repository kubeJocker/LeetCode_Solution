package _653_minimum_deletions_to_make_string_balanced

import "strings"

/*
首先统计字符串中所有a的数量作为起点位置的a的后缀数量
枚举所有字符，如果是‘a’， 则当前位置a的后缀-1，如果是b，当前位置b的前缀+1
最后结果为：a的后缀+b的前缀 的最小值
*/

func minimumDeletions(s string) int {
	countA := strings.Count(s, "a")
	res := countA
	for _, ch := range s {
		if ch == 'a' {
			countA--
		} else {
			countA++
		}
		if countA < res {
			res = countA
		}
	}
	return res
}

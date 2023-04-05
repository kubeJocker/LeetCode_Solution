package _277_find_the_celebrity

func knows(i, j int) bool {
	return true
}

//***********************************************************

func findCelebrity(n int) int {
	// 先假设 cand 是名人
	cand := 0
	for other := 1; other < n; other++ {
		if !knows(other, cand) || knows(cand, other) {
			// cand 不可能是名人，排除
			// 假设 other 是名人
			cand = other
		} else {
			// other 不可能是名人，排除
			// 什么都不用做，继续假设 cand 是名人
		}
	}

	// 现在的 cand 是排除的最后结果，但不能保证一定是名人
	for other := 0; other < n; other++ {
		if cand == other {
			continue
		}
		// 需要保证其他人都认识 cand，且 cand 不认识任何其他人
		if !knows(other, cand) || knows(cand, other) {
			return -1
		}
	}
	return cand
}

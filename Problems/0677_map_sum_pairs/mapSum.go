package _677_map_sum_pairs

type MapSum struct {
	keys map[string]int
}

func Constructor() MapSum {
	return MapSum{make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	this.keys[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	prefixAsRunes, res := []rune(prefix), 0
	for key, val := range this.keys {
		if len(key) >= len(prefix) {
			shouldSum := true
			for i, char := range key {
				if i >= len(prefixAsRunes) {
					break
				}
				if prefixAsRunes[i] != char {
					shouldSum = false
					break
				}
			}
			if shouldSum {
				res += val
			}
		}
	}
	return res
}

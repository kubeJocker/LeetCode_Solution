package _752_open_the_lock

func openLock(deadends []string, target string) int {
	visited := make(map[string]bool)
	for _, deadends := range deadends {
		visited[deadends] = true
	}
	queue := []string{}
	queue = append(queue, "0000")
	step := 0
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			numStr := queue[i]
			if ok := visited[numStr]; ok {
				continue
			}
			if numStr == target {
				return step
			}
			visited[numStr] = true
			for j := 0; j < 4; j++ {
				up := plusOne(numStr, j)
				if ok := visited[up]; !ok {
					queue = append(queue, up)
				}
				down := minuOne(numStr, j)
				if ok := visited[down]; !ok {
					queue = append(queue, down)
				}
			}
		}
		step++
		queue = queue[length:]
	}
	return -1
}

func plusOne(s string, j int) string {
	chars := []byte(s)
	if chars[j] == '9' {
		chars[j] = '0'
	} else {
		chars[j] += 1
	}
	return string(chars)
}

func minuOne(s string, j int) string {
	chars := []byte(s)
	if chars[j] == '0' {
		chars[j] = '9'
	} else {
		chars[j] -= 1
	}
	return string(chars)
}

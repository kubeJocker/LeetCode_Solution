package _042_trapping_rain_water

func trap(height []int) int {
	left, right := 0, len(height)-1
	l_max, r_max, res := 0, 0, 0
	for left < right {
		l_max = max(l_max, height[left])
		r_max = max(r_max, height[right])
		if l_max < r_max {
			res += l_max - height[left]
			left++
		} else {
			res += r_max - height[right]
			right--
		}
	}
	return res
}

func trap2(height []int) int {
	stack, res := []int{}, 0
	for i, h := range height {
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			res += (i - left - 1) * (min(height[left], h) - height[top])
		}
		stack = append(stack, i)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

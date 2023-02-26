package _739_daily_temperatures

func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	res, stack := make([]int, length), []int{}
	for i, num := range temperatures {
		for len(stack) > 0 && num > temperatures[stack[len(stack)-1]] {
			preIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[preIndex] = i - preIndex
		}
		stack = append(stack, i)
	}
	return res
}

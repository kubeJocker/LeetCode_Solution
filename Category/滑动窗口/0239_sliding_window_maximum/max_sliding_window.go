package _239_sliding_window_maximum

func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}
	push := func(i int) {
		for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	for i := 0; i < k; i++ {
		push(i)
	}
	n := len(nums)
	res := make([]int, 1, n-k+1)
	res[0] = nums[queue[0]]
	for i := k; i < n; i++ {
		push(i)
		for queue[0] <= i-k {
			queue = queue[1:]
		}
		res = append(res, nums[queue[0]])
	}
	return res
}

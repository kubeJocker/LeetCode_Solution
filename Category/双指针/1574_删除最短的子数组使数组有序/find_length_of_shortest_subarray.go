package _574_删除最短的子数组使数组有序

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	right := n - 1
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}
	if right == 0 {
		return 0
	}
	res := right
	for i := 0; i < n; i++ {
		for right < n && arr[right] < arr[i] {
			right++
		}
		res = min(res, right-i+1)
		if i+1 < n && arr[i] > arr[i+1] {
			break
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

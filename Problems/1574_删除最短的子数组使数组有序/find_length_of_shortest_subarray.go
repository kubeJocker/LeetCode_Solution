package _574_删除最短的子数组使数组有序

/*
从右遍历一个非递减数列
之后从左遍历，满足左边非递减，右边非递减，且arr[left] < arr[right]
*/
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
	for left := 0; left < n; left++ {
		for right < n && arr[right] < arr[left] {
			right++
		}
		res = min(res, right-left-1)
		if left+1 < n && arr[left] > arr[left+1] {
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

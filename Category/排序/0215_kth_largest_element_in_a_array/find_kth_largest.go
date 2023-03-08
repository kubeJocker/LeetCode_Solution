package _215_kth_largest_element_in_a_array

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSort(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSort(nums []int, left int, right int, index int) int {
	q := randomPartition(nums, left, right)
	if q == index {
		return nums[q]
	} else if q > index {
		return quickSort(nums, left, q-1, index)
	} else {
		return quickSort(nums, q+1, right, index)
	}
}

func randomPartition(nums []int, left int, right int) int {
	i := rand.Int()%(right-left+1) + left
	nums[i], nums[right] = nums[right], nums[i]
	return Partition(nums, left, right)
}

func Partition(nums []int, left int, right int) int {
	temp := nums[left]
	i, j := left+1, right
	for i <= j {
		for i < right && nums[i] <= temp {
			i++
		}
		for j > left && nums[j] > temp {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[j] = nums[j], nums[left]
	return j
}

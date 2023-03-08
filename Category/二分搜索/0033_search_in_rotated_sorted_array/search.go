package _033_search_in_rotated_sorted_array

/*
旋转后必然有一边是顺序，可以根据target是否在顺序区间汇总来取舍区间
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}
	return -1
}

package _034_find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	return []int{left_bound(nums, target), right_bound(nums, target)}
}

func left_bound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if left == len(nums) {
		return -1
	}
	if nums[left] != target {
		return -1
	}
	return left
}

func right_bound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if left-1 < 0 {
		return -1
	}
	if nums[left-1] != target {
		return -1
	}
	return left - 1
}

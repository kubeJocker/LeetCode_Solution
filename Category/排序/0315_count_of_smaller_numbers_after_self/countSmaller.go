package _315_count_of_smaller_numbers_after_self

type Node struct {
	val, index int
}

func countSmaller(nums []int) []int {
	length := len(nums)
	temp, counts, arr := make([]*Node, length), make([]int, length), make([]*Node, length)
	for i := 0; i < length; i++ {
		arr[i] = &Node{nums[i], i}
	}
	sort(arr, 0, length-1, &temp, &counts)
	res := []int{}
	for _, count := range counts {
		res = append(res, count)
	}
	return res
}

func sort(arr []*Node, left int, right int, temp *[]*Node, counts *[]int) {
	if left == right {
		return
	}
	mid := left + (right-left)/2
	sort(arr, left, mid, temp, counts)
	sort(arr, mid+1, right, temp, counts)
	merge(arr, left, right, mid, temp, counts)
}

func merge(arr []*Node, left int, right int, mid int, temp *[]*Node, counts *[]int) {
	for i := left; i <= right; i++ {
		(*temp)[i] = arr[i]
	}
	i, j := left, mid+1
	for p := left; p <= right; p++ {
		if i == mid+1 {
			arr[p] = (*temp)[j]
			j++
		} else if j == right+1 {
			arr[p] = (*temp)[i]
			i++
			(*counts)[arr[p].index] += j - mid - 1
		} else if (*temp)[i].val > (*temp)[j].val {
			arr[p] = (*temp)[j]
			j++
		} else {
			arr[p] = (*temp)[i]
			i++
			(*counts)[arr[p].index] += j - mid - 1
		}
	}
}

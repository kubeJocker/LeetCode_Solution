package _379E_minimim_recolors_to_get_k_consecutive_black_blocks

func minimumRecolors(blocks string, k int) int {
	n := len(blocks)
	if n < k {
		return -1
	}
	left := 0
	curr, min := 0, k
	for right := 0; right < n; right++ {
		if blocks[right] == 'W' {
			curr++
		}
		if right < k-1 {
			continue
		}
		if curr < min {
			min = curr
		}
		if blocks[left] == 'W' {
			curr--
		}
		left++

	}
	return min
}

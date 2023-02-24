package _076_minimum_window_substring

import (
	"math"
)

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	if len(s) < len(t) {
		return ""
	}
	valid := 0
	window, need := map[byte]int{}, map[byte]int{}
	for i := range t {
		need[t[i]]++
	}
	left, right := 0, 0
	start, length := 0, math.MaxInt
	for right < len(s) {
		ch := s[right]
		right++
		if need[ch] > 0 {
			window[ch]++
			if window[ch] == need[ch] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			d := s[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}

		}
	}
	if length == math.MaxInt {
		return ""
	}
	return s[start : start+length]
}

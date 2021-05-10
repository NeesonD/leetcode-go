package _003_lengthOfLongestSubstring

import "math"

func T() {

}

func lengthOfLongestSubstring(s string) int {
	window := make(map[uint8]int)
	var left, right int
	var result int
	for i := range s {
		window[s[i]]++
		right++
		for check(window) {
			window[s[left]]--
			left++
		}
		result = int(math.Max(float64(result), float64(right-left)))
	}

	return result
}

func check(window map[uint8]int) bool {

	for _, i := range window {
		if i > 1 {
			return true
		}
	}
	return false
}

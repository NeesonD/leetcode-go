package _438_findAnagrams

func findAnagrams(s string, p string) []int {
	window := make(map[uint8]int)
	need := make(map[uint8]int)
	for i := range p {
		need[p[i]] ++
	}
	var left, right int
	var result []int
	for i := range s {
		window[s[i]] ++
		right++
		for in(window, need) {
			if right-left == len(p) {
				result = append(result, left)
			}
			window[s[left]] --
			left++
		}
	}
	return result
}

func in(window, need map[uint8]int) bool {
	for k, v := range need {
		i := window[k]
		if i < v {
			return false
		}
	}
	return true
}

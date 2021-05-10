package _076_minWindow

func T() {

}

func minWindow(s string, t string) string {
	window := make(map[uint8]int)
	need := make(map[uint8]int)
	for i := range t {
		need[t[i]]++
	}
	var result string
	var left, right int
	for x := range s {
		window[s[x]]++
		right++
		for in(window, need) {
			if len(result) == 0 || right-left < len(result) {
				result = s[left:right]
			}
			window[s[left]]--
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

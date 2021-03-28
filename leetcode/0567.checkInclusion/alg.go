package _567_checkInclusion

func T()  {
	
}

func checkInclusion(s2 string, s1 string) bool {
	window := make(map[uint8]int)
	need := make(map[uint8]int)
	for i := range s2 {
		need[s2[i]] ++
	}
	var left, right int
	for i := range s1 {
		window[s1[i]]++
		right++
		for in(window,need)   {
			if right - left == len(s2) {
				return true
			}
			window[s1[left]] --
			left ++
		}
	}
	return false
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

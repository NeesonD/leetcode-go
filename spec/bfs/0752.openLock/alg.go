package _752_openLock

func openLock(deadends []string, target string) int {
	deadDict := make(map[string]struct{})
	for _, deadend := range deadends {
		deadDict[deadend] = struct{}{}
		if deadend == "0000" {
			return -1
		}
	}
	queue := make([]string, 0)
	queue = append(queue, "0000")
	visited := make(map[string]struct{})
	visited["0000"] = struct{}{}
	var result int
	for len(queue) != 0 {
		size := len(queue)
		for _, s := range queue {
			if s == target {
				return result
			}
			for i := range s {
				up := up(s, i)
				if _, ok := deadDict[up]; !ok {
					if _, ok := visited[up]; !ok {
						queue = append(queue, up)
					}
				}

				down := down(s, i)
				if _, ok := deadDict[down]; !ok {
					if _, ok := visited[down]; !ok {
						queue = append(queue, down)
					}
				}
				visited[up] = struct{}{}
				visited[down] = struct{}{}
			}
		}
		result++
		queue = queue[size:]
	}
	return -1
}

func up(s string, i int) string {
	u := s[i]
	if u == '9' {
		return s[:i] + "0" + s[i+1:]
	}
	c := s[i] + 1
	return s[:i] + string(c) + s[i+1:]
}

func down(s string, i int) string {
	u := s[i]
	if u == '0' {
		return s[:i] + "9" + s[i+1:]
	}
	return s[:i] + string(s[i]-1) + s[i+1:]
}

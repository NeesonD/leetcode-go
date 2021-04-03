package _051_solveNQueens

func F() {

}

var result [][]string

func solveNQueens(n int) [][]string {
	result = make([][]string, 0)
	record := make([]int, 0)
	var s string
	for i := 0; i < n; i++ {
		s += "."
	}
	for i := 0; i < n; i++ {
		dict[i] = s[:i] + "Q" + s[i+1:]
	}
	help(n, record)
	return result
}

var dict = map[int]string{}

func help(n int, record []int) {
	if len(record) == n {
		tmp := make([]string, 0, len(record))
		for _, s := range record {
			tmp = append(tmp, dict[s])
		}
		result = append(result, tmp)
		return
	}
	for i := 0; i < n; i++ {
		if !check(record, i) {
			continue
		}
		record = append(record, i)
		help(n, record)
		record = record[:len(record)-1]
	}
}

func check(record []int, curr int) bool {
	for i, num := range record {
		if num == curr {
			return false
		}
		if len(record)-i == curr-num || len(record)-i == num-curr {
			return false
		}
	}
	return true
}

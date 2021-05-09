package _077_combine

var result [][]int

func combine(n int, k int) [][]int {
	result = make([][]int, 0)
	record := make([]int, 0)
	help(n, k, record, 0)
	return result
}

func help(n int, k int, record []int, choose int) {
	if len(record) == k {
		tmp := make([]int, k)
		copy(tmp, record)
		result = append(result, tmp)
		return
	}
	for i := 1; i < n+1; i++ {
		if choose >= i {
			continue
		}
		record = append(record, i)
		help(n, k, record, i)
		record = record[:len(record)-1]
	}
}

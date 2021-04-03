package _078_subsets

var result [][]int

func subsets(nums []int) [][]int {
	result = make([][]int, 0)
	record := make([]int, 0)
	help(nums, record, -1)
	return result
}

func help(nums []int, record []int, choose int) {
	tmp := make([]int, len(record))
	copy(tmp, record)
	result = append(result, tmp)
	for i, num := range nums {
		if i <= choose {
			continue
		}
		record = append(record, num)
		help(nums, record, i)
		record = record[:len(record)-1]
	}
}

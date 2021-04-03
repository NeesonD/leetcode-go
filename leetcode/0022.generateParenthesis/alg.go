package _022_generateParenthesis

func F() {

}

var result = make([]string, 0)

func generateParenthesis(n int) []string {
	result = result[0:0]
	var record string
	help(n, record, n, n)
	return result
}

func help(n int, record string, left int, right int) {

	if right < left {
		return
	}
	if left < 0 || right < 0 {
		return
	}
	if len(record) == n*2 {
		result = append(result, record)
		return
	}

	record += string('(')
	help(n, record, left-1, right)
	record = record[:len(record)-1]

	record += string(')')
	help(n, record, left, right-1)
	record = record[:len(record)-1]
}

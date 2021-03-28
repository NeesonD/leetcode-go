package main

import (
	"fmt"
	"math"
)

func main() {
	solveNQueens(4)
}

var item []int

func queue(q int) {
	if len(item) == q {
		fmt.Println(item)
		it := make([]string, 0, q)
		for _, i := range item {
			it = append(it, getString(i, q))
		}
		result = append(result, it)
		return
	}
	for i := 0; i < q; i++ {
		if !ok(item, i) {
			continue
		}
		item = append(item, i)
		queue(q)
		item = item[:len(item)-1]
	}
}

func getString(i int, q int) string {
	var result string
	for j := 0; j < q; j++ {
		if j == i {
			result += "Q"

		} else {
			result += "."
		}

	}
	return result
}

func ok(or []int, curr int) bool {
	for index, point := range or {
		if curr == point {
			return false
		}
		if int(math.Abs(float64(curr-point))) == len(or)-index {
			return false
		}
	}

	return true
}

var result [][]string

func solveNQueens(n int) [][]string {
	result = make([][]string,0)
	queue(n)
	return result
}
